package models

import "time"

type ProjectMetric struct {
    ProjectID    string    `json:"project_id"`
    ServiceName  string    `json:"service_name"`
    MetricType   string    `json:"metric_type"`
    WindowStart  time.Time `json:"window_start"`
    WindowEnd    time.Time `json:"window_end"`
    Value        float64   `json:"value"`
    CreatedAt    time.Time `json:"created_at"`
    CollectedAt  time.Time `json:"collected_at"`
}
