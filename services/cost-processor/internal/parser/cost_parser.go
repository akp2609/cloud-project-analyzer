package parser

import (
	"fmt"
	"strconv"
	"time"

	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/model"
)

func ParserCostRecords(
	records [][]string,
	tenant string,
) ([]model.CostRecord, error) {

	if len(records) < 2{
		return nil, fmt.Errorf("no data rows")
	}

	header := records[0]
	colIndex := map[string]int{}

	for i, h := range header {
		colIndex[h] = i
	}

	required := []string{"service","project","date","cost"}
	for _, r := range required {
		if _, ok := colIndex[r]; !ok {
			return nil, fmt.Errorf("missing columnL %s", r)
		}
	}

	var result []model.CostRecord

	for _, row := range records[1:] {
		cost, err := strconv.ParseFloat(row[colIndex["cost"]],64)
		if err != nil {
			continue
		}

		date, err := time.Parse("2006-01-02", row[colIndex["date"]])
		if err != nil{
			continue
		}

		result = append(result,model.CostRecord{
			Service: row[colIndex["service"]],
			Project: row[colIndex["project"]],
			Date: date,
			Cost: cost,
			Tenant: tenant,
		})
	}

	return result, nil
}