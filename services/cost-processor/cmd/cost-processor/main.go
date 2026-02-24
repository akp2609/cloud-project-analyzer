package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"cloud.google.com/go/bigquery"
	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/aggregator"
	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/model"
	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/parser"
	"github.com/amanpandey1910/cloud-project-analyzer/cost-processor/internal/writer"
)

type PubSubEnvelope struct {
	Message struct{
		Data []byte `json:"data"`
	} `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var env PubSubEnvelope

	if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	
	var event model.UploadEvent
	if err := json.Unmarshal(env.Message.Data, &event); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	log.Printf(
		"Received upload event: bucket=%s path=%s tenant=%s",
		event.Bucket, event.Path, event.Tenant,
	)

	filePath := "/tmp/bucket" + event.Bucket + "/" + event.Path

	records, err := parser.ReadCSV(filePath)
	if err != nil {
		log.Printf("CSV read error: %v", err)
		http.Error(w, "CSV read failed", 500)
		return
	}

	if len(records)<2{
		http.Error(w, "CSV has no data rows", 400)
	}

	log.Printf("CSV row=%d",len(records))
    log.Printf("Header=%v", records[0])
	log.Printf("First row=%v",records[1])

	costRecords, err := parser.ParserCostRecords(records, event.Tenant)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Printf("Parsed %d cost records", len(costRecords))

	aggregated := aggregator.AggregrateByServiceAndDate(costRecords)
	log.Printf("Aggregated %d cost records", len(aggregated))

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		http.Error(w, "gcs client error", 500)
		return
	}
	defer client.Close()

	processedBucket := os.Getenv("PROCESSED_BUCKET")
	if processedBucket == "" {
		http.Error(w, "PROCESSED_BUCKET not set", 500)
		return
	}

	objectPath := fmt.Sprintf(
		"tenant=%s/date=%s/aggregated.json",
		event.Tenant,
		aggregated[0].Date.Format("2006-01-02"),
	)

	err = writer.WriteAggredatedJSONToGCS(
		ctx,
		client,
		processedBucket,
		objectPath,
		aggregated,
	)

	if err != nil {
		log.Printf("GCS write failed: %v", err)
		http.Error(w, "gcs write failed", 500)
		return
	}

	log.Printf("Processed data written to gs://%s/%s", processedBucket, objectPath)

	bqClient, err := bigquery.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		http.Error(w, "bq client error", 500)
		return
	}
	defer bqClient.Close()

	err = writer.InsertAggregatedCost(
		ctx,
		bqClient,
		"cost_analytics",
		"processed_costs",
		aggregated,
	)

	if err != nil {
		log.Printf("BigQuery insert failed: %v",err)
		http.Error(w, "bq insert failed", 500)
		return
	}

	log.Printf("Inserted %d rows into Bigquery", len(aggregated))

	w.WriteHeader(http.StatusOK)
}


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/",handler)
	mux.HandleFunc("/livez", func(w http.ResponseWriter, _ *http.Request){
		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/readyz", func (w http.ResponseWriter, _ *http.Request)  {
		w.Write([]byte("OK"))
	})

	log.Printf("Cost processor listening on %s", port)
	log.Fatal(http.ListenAndServe(":"+port,mux))
}