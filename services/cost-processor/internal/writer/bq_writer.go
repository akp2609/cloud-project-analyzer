package writer

import (
	"context"
	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/model"
	"cloud.google.com/go/bigquery"
)


func InsertAggregatedCost(
	ctx context.Context,
	client *bigquery.Client,
	dataset string,
	table string,
	data []model.AggregatedCost,
) error {

	var rows []*model.BqRow

	for _, r := range data {
		rows = append(rows, &model.BqRow{
			Tenant: r.Tenant,
			Service: r.Service,
			Date: r.Date.Format("2006-01-02"),
			TotalCost: r.Total,
		})
	}

	inserter := client.Dataset(dataset).Table(table).Inserter()
	return inserter.Put(ctx,rows)
}
