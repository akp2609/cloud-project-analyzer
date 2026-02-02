package ingestion

import (
	"context"
	"database/sql"
	"log"
	"time"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/civil"
	"google.golang.org/api/iterator"
)

type CostRow struct {
	ProjectID string    `bigquery:"tenant"`
	Service   string    `bigquery:"service"`
	UsageDate civil.Date `bigquery:"date"`
	Cost      float64   `bigquery:"total_cost"`
}

type CostSyncer struct {
	bq      *bigquery.Client
	db      *sql.DB
	dataset string
	table   string
}

func NewCostSyncer(bq *bigquery.Client, db *sql.DB, dataset, table string) *CostSyncer {
	return &CostSyncer{
		bq:      bq,
		db:      db,
		dataset: dataset,
		table:   table,
	}
}

func (s *CostSyncer) lastSyncedDate(
	ctx context.Context,
	projectID string,
) (time.Time, error) {

	var lastDate sql.NullTime

	err := s.db.QueryRowContext(
		ctx,
		`SELECT MAX(usage_date) FROM cost_daily WHERE project_id = $1`,
		projectID,
	).Scan(&lastDate)

	if err != nil {
		return time.Time{}, err
	}

	if !lastDate.Valid {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), nil
	}

	return lastDate.Time, nil
}

func (s *CostSyncer) SyncProject(ctx context.Context, projectID string) error {
	lastDate, err := s.lastSyncedDate(ctx, projectID)
	if err != nil {
		return err
	}

	query := s.bq.Query(`
		SELECT tenant, service, date, total_cost
		FROM ` + "`" + s.dataset + "." + s.table + "`" + `
		WHERE tenant = @project_id
		  AND date > DATE(@last_date)
	`)

	query.Parameters = []bigquery.QueryParameter{
		{Name: "project_id", Value: projectID},
		{Name: "last_date", Value: lastDate},
	}

	it, err := query.Read(ctx)
	if err != nil {
		return err
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx, `
		INSERT INTO cost_daily (project_id, service, cost, usage_date)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (project_id, service, usage_date)
		DO UPDATE SET cost = EXCLUDED.cost
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for {
		var row CostRow
		err := it.Next(&row)
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}

		_, err = stmt.ExecContext(
			ctx,
			row.ProjectID,
			row.Service,
			row.Cost,
			row.UsageDate,
		)
		if err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	log.Printf("Cost sync completed for project %s", projectID)
	return nil
}
