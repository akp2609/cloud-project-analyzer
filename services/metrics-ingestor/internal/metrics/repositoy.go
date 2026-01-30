package metrics

import (
	"context"
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) UpsertMetric(
    ctx context.Context,
    projectID, serviceName, metricType string,
    windowStart, windowEnd time.Time,
    value float64,
    collectedAt time.Time,
) error {
    query := `
    INSERT INTO project_metrics (
        project_id, service_name, metric_type, window_start, window_end,
        value, collected_at
    )
    VALUES ($1, $2, $3, $4, $5, $6, $7)
    ON CONFLICT (project_id, service_name, metric_type, window_start)
    DO UPDATE SET
        value = EXCLUDED.value,
        collected_at = EXCLUDED.collected_at;
    `
    _, err := r.db.ExecContext(ctx, query,
        projectID, serviceName, metricType,
        windowStart, windowEnd,
        value, collectedAt,
    )
    return err
}
