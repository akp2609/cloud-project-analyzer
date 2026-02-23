package models

import "time"

type LinkedProject struct {
    ID                 string    `json:"id"`
    ProjectID          string    `json:"project_id"`
    ProjectNumber      string    `json:"project_number"`
    BillingAccount     string    `json:"billing_account"`
    Status             string    `json:"status"`
    IAMVerified        bool      `json:"iam_verified"`
    BillingVerified    bool      `json:"billing_verified"`
    MonitoringVerified bool      `json:"monitoring_verified"`
    LastVerifiedAt     *time.Time `json:"last_verified_at,omitempty"`
    ErrorReason        string    `json:"error_reason,omitempty"`
    CreatedAt          time.Time `json:"created_at"`
    UpdatedAt          time.Time `json:"updated_at"`
}
