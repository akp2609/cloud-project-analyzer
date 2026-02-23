package repository

import (
    "context"
    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/models"
)

func (r *Repository) GetAllProjects(ctx context.Context) ([]models.LinkedProject, error) {
    rows, err := r.db.QueryContext(ctx, `
        SELECT id, project_id, project_number, billing_account, status,
               iam_verified, billing_verified, monitoring_verified,
               last_verfied_at, error_reason, created_at, updated_at
        FROM linked_projects
        ORDER BY created_at ASC
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []models.LinkedProject
    for rows.Next() {
        var p models.LinkedProject
        err := rows.Scan(
            &p.ID,
            &p.ProjectID,
            &p.ProjectNumber,
            &p.BillingAccount,
            &p.Status,
            &p.IAMVerified,
            &p.BillingVerified,
            &p.MonitoringVerified,
            &p.LastVerifiedAt,
            &p.ErrorReason,
            &p.CreatedAt,
            &p.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        projects = append(projects, p)
    }

    return projects, nil
}
