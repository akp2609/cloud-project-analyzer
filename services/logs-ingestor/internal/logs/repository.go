package logs

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

func (r *Repository) IncrementError(
	ctx context.Context,
	projectID string,
	at time.Time,
) error {

	query := `
	INSERT INTO project_log_stats (
		project_id,
		error_count,
		last_error_at
	)
	VALUES ($1, 1, $2)
	ON CONFLICT (project_id)
	DO UPDATE SET
		error_count = project_log_stats.error_count + 1,
		last_error_at = EXCLUDED.last_error_at,
		updated_at = now();
	`

	_, err := r.db.ExecContext(ctx, query, projectID, at)
	return err
}
