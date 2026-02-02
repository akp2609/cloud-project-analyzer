package models

import (
    "time"

    "github.com/google/uuid"
    "encoding/json"
)

type ProjectInsight struct {
    ID          uuid.UUID       `db:"id"`           
    ProjectID   string          `db:"project_id"`   
    InsightType string          `db:"insight_type"` 
    Severity    string          `db:"severity"`     
    Title       string          `db:"title"`
    Description string          `db:"description"`
    DetectedAt  time.Time       `db:"detected_at"`
    Metadata    json.RawMessage `db:"metadata"`     
}
