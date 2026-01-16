package aggregator

import (
	"time"

	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/model"
)

func AggregrateByServiceAndDate(
	records []model.CostRecord,
) []model.AggregatedCost {

	type key struct {
		service string
		date time.Time
		tenant string
	}

	m := make(map[key]float64)

	for _, r := range records {
		k := key{
			service: r.Service,
			date: r.Date,
			tenant: r.Tenant,
		}
		m[k] += r.Cost
	}

	var result []model.AggregatedCost
	for k, total := range m {
		result = append(result, model.AggregatedCost{
			Service: k.service,
			Date: k.date,
			Total: total,
			Tenant: k.tenant,
		})
	}

	return result
} 