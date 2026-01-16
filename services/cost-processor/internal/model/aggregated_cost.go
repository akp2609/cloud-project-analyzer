package model

import "time"

type AggregatedCost struct {
	Service string
	Date time.Time
	Total float64
	Tenant string
}