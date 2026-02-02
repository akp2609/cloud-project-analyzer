package api

import (
    "context"
    "encoding/json"
    "net/http"
    "log"

    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/analysis"
    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/repository"
)

type Handler struct {
    repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
    return &Handler{repo: repo}
}

func (h *Handler) GetProjectCostAnomalies(w http.ResponseWriter, r *http.Request) {
    projectID := r.URL.Query().Get("project_id")
    if projectID == "" {
        http.Error(w, "project_id required", http.StatusBadRequest)
        return
    }

    ctx := context.Background()

    records, err := h.repo.GetDailyCosts(ctx, projectID)
    if err != nil {
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }

    anomalies := analysis.DetectCostAnomalies(records)

    for _, a := range anomalies {
	err := h.repo.InsertCostAnomaly(ctx, a)
	if err != nil {
		log.Println("cost anomaly insert failed:", err)
		continue
	}

	insight := analysis.BuildCostSpikeInsight(a)
	_ = h.repo.InsertProjectInsight(ctx, insight)
    }


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(anomalies)
}

func (h *Handler) Routes() http.Handler {
    mux := http.NewServeMux()
    mux.HandleFunc("/projects/anomalies", h.GetProjectCostAnomalies)
    return mux
}
