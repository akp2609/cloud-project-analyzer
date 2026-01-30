package logs

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) HandleMessage(ctx context.Context, data []byte) error {
	var entry LogEntry

	if err := json.Unmarshal(data, &entry); err != nil {
		return err
	}

	projectID := entry.Resource.Labels.ProjectID
	if projectID == "" {
		return nil
	}

	severity := strings.ToUpper(entry.Severity)
	if severity != "ERROR" && severity != "CRITICAL" && severity != "ALERT" {
		return nil 
	}

	timestamp, err := time.Parse(time.RFC3339, entry.Timestamp)
	if err != nil {
		timestamp = time.Now()
	}

	log.Printf("Error log detected for project=%s", projectID)

	return h.repo.IncrementError(ctx, projectID, timestamp)
}
