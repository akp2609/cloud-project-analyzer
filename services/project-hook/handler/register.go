package handler

import (
	"log"
	"database/sql"
	"encoding/json"
	"net/http"
)

type RegisterProjectRequest struct {
	ProjectID string `json:"project_id"`
}

func RegisterProject(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RegisterProjectRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", 400)
			return
		}

		if req.ProjectID == "" {
			http.Error(w, "project_id required", 400)
			return
		}

		_, err := db.Exec(`
			INSERT INTO linked_projects (
				project_id,
				project_number,
				status
			)
			VALUES ($1, $2, 'PENDING')
			ON CONFLICT (project_id) DO NOTHING
		`, req.ProjectID, "unknown")

		if err != nil {
	    log.Printf("DB insert failed: %v", err)
	    http.Error(w, err.Error(), 500)
	    return
   }


		w.WriteHeader(http.StatusCreated)
	}
}
