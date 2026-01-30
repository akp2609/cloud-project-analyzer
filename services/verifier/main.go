package main

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

type PubSubEnvelope struct {
	Message struct {
		Data string `json:"data"`
	} `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request){
	var env PubSubEnvelope
	if err := json.NewDecoder(r.Body).Decode(&env); err != nil {
		http.Error(w, "bad request", 400)
		return
	}

	raw, err := base64.StdEncoding.DecodeString(string(env.Message.Data))
	if err != nil {
		http.Error(w, "invalid base64",400)
		return
	}

	var event ProjectLinkedEvent
	if err := json.Unmarshal(raw,&event); err != nil {
		http.Error(w, "invalid payload", 400)
		return
	}

	log.Printf("Verifier received project: %s", event.ProjectID)

	ctx := r.Context()

	if err := verifyProject(ctx, event.ProjectID); err != nil {
		log.Printf("Project verification failed: %v", err)
		http.Error(w,"Project not accessible", http.StatusForbidden)
		return
	}

	log.Printf("Project verified: %s", event.ProjectID)

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := verifyProject(ctx, event.ProjectID); err != nil {
		_ = markFailed(ctx, db, event.ProjectID, err.Error())
		http.Error(w, "project not accessible", http.StatusForbidden)
		return
	}

	if err := markVerified(ctx, db, event.ProjectID); err != nil { 
		log.Printf("markVerified failed: %v", err)
		http.Error(w, "db update failed", 500) 
		return 
	}

	w.WriteHeader(200)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/livez", func(w http.ResponseWriter,_ *http.Request){
		w.Write([]byte("OK"))
	})

	log.Println("Verifier listening on :8080")
	log.Fatal(http.ListenAndServe(":8080",mux))
}