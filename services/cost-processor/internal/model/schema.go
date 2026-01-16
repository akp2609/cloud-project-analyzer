package model

type BqRow struct {
	Tenant string `bigquery:"tenant"`
	Service string `bigquery:"service"`
	Date string `bigquery:"date"`
	TotalCost float64 `bigquery:"total_cost"`
}