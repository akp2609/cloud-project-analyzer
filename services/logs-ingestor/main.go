package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/amanpandey1910/cloud-project-analyzer/services/logs-ingestor/internal/logs"
)

type PubSubEnvelope struct {
	Message struct {
		Data string `json:"data"`
	} `json:"message"`
}

func mustOpenDB() *sql.DB {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open db: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	return db
}


func handler(w http.ResponseWriter, r *http.Request, logHandler *logs.Handler) {
	
	ctx := r.Context()

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

		if err := logHandler.HandleMessage(ctx, raw); err != nil {
			log.Printf("failed to handle log: %v", err)
			http.Error(w, "processing failed", 500)
			return
		}

		w.WriteHeader(http.StatusOK)
}

func main() {
    db := mustOpenDB()
    repo := logs.NewRepository(db)
    logHandler := logs.NewHandler(repo)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        handler(w, r, logHandler)
    })

    log.Println("Logs ingestor listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}


