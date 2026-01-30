package model

import "time"

type LinkedProject struct {
	ID string
	ProjectID string
	ProjectNumber string
	BillingAccount *string
	Status string

	IAMVerified bool
	BillingVerified bool
	MonitoringVerified bool

	LastVerifiedAt *time.Time
	ErrorReason *string

	CreatedAt time.Time
	UpdatedAt time.Time
}