package analysis

import (
    "context"
    "math"

    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/models"
)

type MetricInsight struct {
    ProjectID  string  `json:"project_id"`
    MetricType string  `json:"metric_type"`
    Avg        float64 `json:"avg"`
    Max        float64 `json:"max"`
    Min        float64 `json:"min"`
    Count      int     `json:"count"`
}

func CalculateInsights(ctx context.Context, metrics []models.ProjectMetric) MetricInsight {
    if len(metrics) == 0 {
        return MetricInsight{}
    }
    var sum, max, min float64
    min = math.MaxFloat64
    for _, m := range metrics {
        sum += m.Value
        if m.Value > max {
            max = m.Value
        }
        if m.Value < min {
            min = m.Value
        }
    }
    return MetricInsight{
        ProjectID:  metrics[0].ProjectID,
        MetricType: metrics[0].MetricType,
        Avg:        sum / float64(len(metrics)),
        Max:        max,
        Min:        min,
        Count:      len(metrics),
    }
}
