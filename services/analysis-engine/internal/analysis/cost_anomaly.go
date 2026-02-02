package analysis

import (
	"github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/models"
)



func DetectCostAnomalies(records []models.CostRecord) []models.CostAnomaly {
	var anomalies []models.CostAnomaly

	if len(records) < 2 {
		return anomalies
	}

	var sum float64
	for i := 0; i < len(records)-1; i++ {
		sum += records[i].Cost
	}

	avg := sum / float64(len(records)-1)
	latest := records[len(records)-1]

	deviation := ((latest.Cost - avg) / avg) * 100

    if latest.Cost > avg*1.5 {
	anomalies = append(anomalies, models.CostAnomaly{
		ProjectID:    latest.ProjectID,
		Service:      latest.Service,
		Date:         latest.Date,

		BaselineCost: avg,
		Cost:   latest.Cost,
		SpikeRatio:   deviation,
		Severity:     SeverityFromDeviation(deviation),
	  })
    }


	return anomalies
}
