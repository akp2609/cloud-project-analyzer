package models

import "time"

type CostRecord struct {
	ProjectID string
	Service   string
	Cost      float64
	Date      time.Time
}
