package models

import "time"

type CostAnomaly struct {
	ProjectID    string
	Service      string
	Date         time.Time 
	Cost         float64
	BaselineCost float64
	SpikeRatio   float64
	Severity     string
}
