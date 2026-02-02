package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
	"cloud.google.com/go/bigquery"

	"github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/api"
	"github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/ingestion"
	"github.com/amanpandey1910/cloud-project-analyzer/analysis-engine/internal/repository"
)

func main() {
	/* -------------------- ENV -------------------- */

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	gcpProjectID := os.Getenv("GCP_PROJECT_ID")
	if gcpProjectID == "" {
		log.Fatal("GCP_PROJECT_ID not set")
	}

	bqDataset := os.Getenv("BQ_DATASET")
	if bqDataset == "" {
		log.Fatal("BQ_DATASET not set")
	}

	bqTable := os.Getenv("BQ_TABLE")
	if bqTable == "" {
		log.Fatal("BQ_TABLE not set")
	}

	targetProjectID := os.Getenv("TARGET_PROJECT_ID")
	if targetProjectID == "" {
		log.Fatal("TARGET_PROJECT_ID not set (tenant/project to sync)")
	}

	/* -------------------- POSTGRES -------------------- */

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatal("Postgres not reachable:", err)
	}

	/* -------------------- BIGQUERY -------------------- */

	ctx := context.Background()

	bqClient, err := bigquery.NewClient(ctx, gcpProjectID)
	if err != nil {
		log.Fatal("Failed to create BigQuery client:", err)
	}
	defer bqClient.Close()

	/* -------------------- COST SYNC -------------------- */

	log.Println("Starting BigQuery → Postgres cost sync...")

	costSyncer := ingestion.NewCostSyncer(
		bqClient,
		db,
		bqDataset,
		bqTable,
	)

	if err := costSyncer.SyncProject(ctx, targetProjectID); err != nil {
		log.Fatal("Cost sync failed:", err)
	}

	log.Println("Cost sync completed successfully")

	/* -------------------- HTTP SERVER -------------------- */

	repo := repository.NewRepository(db)
    handler := api.NewHandler(repo)

    log.Println("Analysis engine listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", handler.Routes()))


}