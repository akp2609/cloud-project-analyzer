package model

import "time"

type CostRecord struct{
	Service string
	Project string
	Date time.Time
	Cost float64
	Tenant string
}