package main

import (
	"log"
	"net/http"
        "fmt"
        "os"

	"github.com/amanpandey1910/cloud-project-analyzer/services/project-hook/internal/db"
	"github.com/amanpandey1910/cloud-project-analyzer/services/project-hook/handler" 
)

func main() {
        fmt.Println("DATABASE_URL =", os.Getenv("DATABASE_URL"))
	database, err := db.NewPostgres()
	if err != nil {
		log.Fatalf("db connection failed: %v", err)
	}
	defer database.Close()

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/projects/link",
		handler.RegisterProject(database),
	)

	mux.HandleFunc("/livez", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("OK"))
	})

	log.Println("Project hook service listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
