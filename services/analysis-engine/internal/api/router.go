package api

import (
    "context"
    "encoding/json"
    "net/http"
    "log"
    "time"

    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/analysis"
    "github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/repository"
)

type Handler struct {
    repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
    return &Handler{repo: repo}
}

func (h *Handler) GetProjectCostAnomalies(w http.ResponseWriter, r *http.Request) {
    projectID := r.URL.Query().Get("project_id")
    if projectID == "" {
        http.Error(w, "project_id required", http.StatusBadRequest)
        return
    }

    ctx := context.Background()

    records, err := h.repo.GetDailyCosts(ctx, projectID)
    if err != nil {
        log.Println("dashboard query failed:", err)
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }

    anomalies := analysis.DetectCostAnomalies(records)

    for _, a := range anomalies {
	err := h.repo.InsertCostAnomaly(ctx, a)
	if err != nil {
		log.Println("cost anomaly insert failed:", err)
		continue
	}

	insight := analysis.BuildCostSpikeInsight(a)
	_ = h.repo.InsertProjectInsight(ctx, insight)
    }


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(anomalies)
}


func (h *Handler) GetDashboardSummary(w http.ResponseWriter, r *http.Request) {
	projectID := r.URL.Query().Get("project_id")
	if projectID == "" {
		http.Error(w, "project_id required", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	summary, err := h.repo.GetDashboardSummary(ctx, projectID)
	if err != nil {
		log.Println("dashboard query failed:", err)
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(summary)
}

func (h *Handler) GetProjectInsights(w http.ResponseWriter, r *http.Request) {
    projectID := r.URL.Query().Get("project_id")
    if projectID == "" {
        http.Error(w, "project_id required", http.StatusBadRequest)
        return
    }

    ctx := context.Background()
    insights, err := h.repo.GetProjectInsights(ctx, projectID)
    if err != nil {
        log.Println("insights query failed:", err)
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(insights)
}

func (h *Handler) GetProjects(w http.ResponseWriter, r *http.Request){
    ctx := context.Background()
    projects, err := h.repo.GetAllProjects(ctx)
    if err != nil {
        log.Println("projects query failed",err)
        http.Error(w, "db error", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json") 
    json.NewEncoder(w).Encode(projects)
}

func (h *Handler) GetMetricsInsights(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    projectID := r.URL.Query().Get("project_id")
    metricType := r.URL.Query().Get("metric_type")

    metrics, err := h.repo.GetMetrics(ctx, projectID, metricType, time.Now().Add(-30*24*time.Hour))
    if err != nil {
        http.Error(w, "db error", 500)
        return
    }

    insight := analysis.CalculateInsights(ctx, metrics)
    insight.ProjectID = projectID
    insight.MetricType = metricType

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(insight)
}



func (h *Handler) Routes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/dashboard", h.GetDashboardSummary)
    mux.HandleFunc("/projects/anomalies", h.GetProjectCostAnomalies)
    mux.HandleFunc("/projects/insights", h.GetProjectInsights)
    mux.HandleFunc("/projects/all", h.GetProjects)
    mux.HandleFunc("/projects/metrics-insight",h.GetMetricsInsights)
    return mux
}

