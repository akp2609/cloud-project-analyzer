package analysis

import (
    "encoding/json"
    "fmt"

    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/models"
)

// Todo more insights
func BuildCostSpikeInsight(a models.CostAnomaly) models.ProjectInsight {
    
    metadataMap := map[string]any{
        "service":       a.Service,
        "date":          a.Date.Format("2006-01-02"), 
        "cost":          a.Cost,
        "baseline_cost": a.BaselineCost,
        "spike_ratio":   a.SpikeRatio,
    }

    
    metadataBytes, _ := json.Marshal(metadataMap)

    return models.ProjectInsight{
        ProjectID:   a.ProjectID,
        InsightType: "COST_SPIKE",
        Severity:    a.Severity,
        Title:       fmt.Sprintf("%s cost spike detected", a.Service),
        Description: fmt.Sprintf(
            "%s cost increased by %.1f%% compared to historical baseline on %s.",
            a.Service,
            a.SpikeRatio,
            a.Date.Format("2006-01-02"),
        ),
        Metadata: metadataBytes, 
    }
}
