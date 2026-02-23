package models

import "time"
import "database/sql"

type LinkedProject struct {
    ID                 string         `json:"id"`
    ProjectID          string         `json:"project_id"`
    ProjectNumber      sql.NullString `json:"project_number"`
    BillingAccount     sql.NullString `json:"billing_account"`
    Status             string         `json:"status"`
    IAMVerified        bool           `json:"iam_verified"`
    BillingVerified    bool           `json:"billing_verified"`
    MonitoringVerified bool           `json:"monitoring_verified"`
    LastVerifiedAt     sql.NullTime   `json:"last_verified_at"`
    ErrorReason        sql.NullString `json:"error_reason"`
    CreatedAt          time.Time      `json:"created_at"`
    UpdatedAt          time.Time      `json:"updated_at"`
}
