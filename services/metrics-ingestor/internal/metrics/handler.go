package metrics

import (
    "context"
    "fmt"
    "log"
    "time"

    monitoring "cloud.google.com/go/monitoring/apiv3/v2"
    monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
    "google.golang.org/api/iterator"
    "google.golang.org/protobuf/types/known/durationpb"
    "google.golang.org/protobuf/types/known/timestamppb"
)


type Handler struct {
    repo *Repository
}

func NewHandler(repo *Repository) *Handler {
    return &Handler{repo: repo}
}

func (h *Handler) Handle(ctx context.Context, projectID string) error {
    client, err := monitoring.NewMetricClient(ctx)
    if err != nil {
        return err
    }
    defer client.Close()

    
    var cloudRunMetrics = []struct {
    Type    string
    Aligner monitoringpb.Aggregation_Aligner
    }{
    {"run.googleapis.com/request_count", monitoringpb.Aggregation_ALIGN_SUM},
    {"run.googleapis.com/error_count", monitoringpb.Aggregation_ALIGN_SUM},
    {"run.googleapis.com/request_latencies", monitoringpb.Aggregation_ALIGN_PERCENTILE_95},
    {"run.googleapis.com/container/cpu/utilization", monitoringpb.Aggregation_ALIGN_MEAN},
    {"run.googleapis.com/container/memory/utilization", monitoringpb.Aggregation_ALIGN_MEAN},
    {"run.googleapis.com/container/restart_count", monitoringpb.Aggregation_ALIGN_SUM},
    {"run.googleapis.com/container/concurrent_requests", monitoringpb.Aggregation_ALIGN_MEAN},
    {"run.googleapis.com/container/instance_count", monitoringpb.Aggregation_ALIGN_MEAN},
    }


    
    for _, m := range cloudRunMetrics {
        req := &monitoringpb.ListTimeSeriesRequest{
            Name:   "projects/" + projectID,
            Filter: fmt.Sprintf(`metric.type="%s"`, m.Type),
            Interval: &monitoringpb.TimeInterval{
                EndTime:   timestamppb.New(time.Now().UTC()),
                StartTime: timestamppb.New(time.Now().Add(-24 * time.Hour).UTC()),
            },
            View: monitoringpb.ListTimeSeriesRequest_FULL,
            Aggregation: &monitoringpb.Aggregation{
                AlignmentPeriod:  durationpb.New(time.Hour),
                PerSeriesAligner: m.Aligner,
            },
        }

        it := client.ListTimeSeries(ctx, req)

        var total float64
        var count int

        for {
            ts, err := it.Next()
            if err == iterator.Done {
                break
            }
            if err != nil {
                log.Printf("metrics error for %s: %v", m.Type, err)
                break
            }

            for _, p := range ts.Points {
                val := p.Value.GetDoubleValue()
                total += val
                count++

                
                if err := h.repo.UpsertMetric(
                    ctx,
                    projectID,
                    "cloud_run",
                    m.Type,
                    p.Interval.StartTime.AsTime(),
                    p.Interval.EndTime.AsTime(),
                    val,
                    time.Now(),
                ); err != nil {
                    log.Printf("db error for %s: %v", m.Type, err)
                }
            }
        }

        if count > 0 {
            avg := total / float64(count)
            log.Printf("project=%s metric=%s avg=%.2f", projectID, m.Type, avg)
        } else {
            log.Printf("no data for metric=%s project=%s", m.Type, projectID)
        }
    }

    return nil
}
