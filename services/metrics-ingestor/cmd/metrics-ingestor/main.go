package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"github.com/amanpandey1910/cloud-project-analyzer/services/metrics-ingestor/internal/metrics"
)

type PubSubEnvelope struct {
	Message struct {
		Data string `json:"data"`
	} `json:"message"`
}

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := metrics.NewRepository(db)
	handler := metrics.NewHandler(repo)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		var env PubSubEnvelope
		if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
			http.Error(w, "bad request", 400)
			return
		}

		raw, err := base64.StdEncoding.DecodeString(env.Message.Data)
		if err != nil {
			http.Error(w, "invalid base64", 400)
			return
		}

		var event metrics.MetricsEvent
		if err := json.Unmarshal(raw, &event); err != nil {
			http.Error(w, "invalid payload", 400)
			return
		}

		if err := handler.Handle(ctx, event.ProjectID); err != nil {
			log.Printf("metrics error: %v", err)
			http.Error(w, "processing failed", 500)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	log.Println("Metrics ingestor listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
