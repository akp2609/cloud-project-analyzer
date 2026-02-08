package repository

import (
    "context"
    "database/sql"
    "time" 

    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/models"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) GetDailyCosts(
    ctx context.Context,
    projectID string,
) ([]models.CostRecord, error) {

    rows, err := r.db.QueryContext(ctx, `
        SELECT project_id, service, cost, usage_date
        FROM cost_daily
        WHERE project_id = $1
        ORDER BY usage_date ASC
    `, projectID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var records []models.CostRecord
    for rows.Next() {
        var rec models.CostRecord
        if err := rows.Scan(
            &rec.ProjectID,
            &rec.Service,
            &rec.Cost,
            &rec.Date,
        ); err != nil {
            return nil, err
        }
        records = append(records, rec)
    }

    return records, nil
}

type DashboardSummary struct {
	ProjectID            string  `json:"project_id"`
	TotalCost30d         float64 `json:"total_cost_30d"`
	ActiveAnomalies      int     `json:"active_anomalies"`
	HighSeverityInsights int     `json:"high_severity_insights"`
}

func (r *Repository) GetDashboardSummary(
	ctx context.Context,
	projectID string,
) (*DashboardSummary, error) {

	var summary DashboardSummary
	summary.ProjectID = projectID

	err := r.db.QueryRowContext(ctx, `
		SELECT COALESCE(SUM(cost), 0)
		FROM cost_daily
		WHERE project_id = $1
		AND usage_date >= CURRENT_DATE - INTERVAL '30 days'
	`, projectID).Scan(&summary.TotalCost30d)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM cost_anomalies
		WHERE project_id = $1
	`, projectID).Scan(&summary.ActiveAnomalies)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRowContext(ctx, `
		SELECT COUNT(*)
		FROM project_insights
		WHERE project_id = $1
		AND severity = 'HIGH'
	`, projectID).Scan(&summary.HighSeverityInsights)
	if err != nil {
		return nil, err
	}

	return &summary, nil
}


func (r *Repository) InsertCostAnomaly(ctx context.Context, a models.CostAnomaly) error {
    query := `
        INSERT INTO cost_anomalies (
            project_id,
            service,
            date,
            cost,
            baseline_cost,
            spike_ratio,
            severity
        )
        VALUES ($1,$2,$3,$4,$5,$6,$7)
        ON CONFLICT (project_id, service, date) DO NOTHING
    `

    dateOnly := a.Date.UTC().Truncate(24 * time.Hour)

    _, err := r.db.ExecContext(
        ctx,
        query,
        a.ProjectID,
        a.Service,
        dateOnly,
        a.Cost,
        a.BaselineCost,
        a.SpikeRatio,
        a.Severity,
    )

    return err
}


func (r *Repository) GetCostAnomaliesByProject(
	ctx context.Context,
	projectID string,
) ([]models.CostAnomaly, error) {

	query := `
	SELECT
		project_id,
		service,
		date,
		cost,
		baseline_cost,
		spike_ratio,
		severity,
		created_at
	FROM cost_anomalies
	WHERE project_id = $1
	ORDER BY date DESC
	`

	rows, err := r.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var anomalies []models.CostAnomaly

	for rows.Next() {
		var a models.CostAnomaly
		err := rows.Scan(
			&a.ProjectID,
			&a.Service,
			&a.Date,
			&a.Cost,
			&a.BaselineCost,
			&a.SpikeRatio,
			&a.Severity,
		)
		if err != nil {
			return nil, err
		}
		anomalies = append(anomalies, a)
	}

	return anomalies, nil
}
