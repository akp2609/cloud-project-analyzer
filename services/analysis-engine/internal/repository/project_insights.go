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
