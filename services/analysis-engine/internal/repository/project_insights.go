package repository

import (
	"context"
	"github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/models"
)

func (r *Repository) InsertProjectInsight(
	ctx context.Context,
	i models.ProjectInsight,
) error {

	query := `
	INSERT INTO project_insights (
		project_id,
		insight_type,
		severity,
		title,
		description,
		detected_at,
		metadeta
	)
	VALUES ($1,$2,$3,$4,$5,now(),$6)
	ON CONFLICT DO NOTHING
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		i.ProjectID,
		i.InsightType,
		i.Severity,
		i.Title,
		i.Description,
		i.Metadata,
	)

	return err
}


func (r *Repository) GetProjectInsights(ctx context.Context, projectID string) ([]models.ProjectInsight, error) {
    rows, err := r.db.QueryContext(ctx, `
        SELECT project_id, insight_type, severity, title, description, detected_at, metadeta
        FROM project_insights
        WHERE project_id = $1
        ORDER BY detected_at DESC
    `, projectID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var insights []models.ProjectInsight
    for rows.Next() {
        var i models.ProjectInsight
        if err := rows.Scan(
            &i.ProjectID,
            &i.InsightType,
            &i.Severity,
            &i.Title,
            &i.Description,
            &i.DetectedAt,
            &i.Metadata,
        ); err != nil {
            return nil, err
        }
        insights = append(insights, i)
    }
    return insights, nil
}
