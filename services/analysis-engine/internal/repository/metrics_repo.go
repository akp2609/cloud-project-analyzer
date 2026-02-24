package repository

import (
    "context"
    "time"
	"database/sql"

    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/models"
)

func (r *Repository) GetMetrics(ctx context.Context, projectID, metricType string, since time.Time) ([]models.ProjectMetric, error) {
    var rows *sql.Rows
    var err error

    if metricType == "" {
        rows, err = r.db.QueryContext(ctx, `
            SELECT project_id, service_name, metric_type, window_start, window_end, value, created_at, collected_at
            FROM project_metrics
            WHERE project_id = $1 AND window_start >= $2
            ORDER BY window_start DESC
        `, projectID, since)
    } else {
        rows, err = r.db.QueryContext(ctx, `
            SELECT project_id, service_name, metric_type, window_start, window_end, value, created_at, collected_at
            FROM project_metrics
            WHERE project_id = $1 AND metric_type = $2 AND window_start >= $3
            ORDER BY window_start DESC
        `, projectID, metricType, since)
    }
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var metrics []models.ProjectMetric
    for rows.Next() {
        var m models.ProjectMetric
        if err := rows.Scan(&m.ProjectID, &m.ServiceName, &m.MetricType, &m.WindowStart, &m.WindowEnd, &m.Value, &m.CreatedAt, &m.CollectedAt); err != nil {
            return nil, err
        }
        metrics = append(metrics, m)
    }
    return metrics, nil
}
