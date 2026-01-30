package main

import (
	"context"
	"database/sql"
	"time"
	"log"
)

func markVerified(ctx context.Context, db *sql.DB, projectID string) error {
	res, err := db.ExecContext(ctx, `
	UPDATE linked_projects
	SET status = 'ACTIVE',
	    last_verfied_at = $2,
		error_reason = NULL
	WHERE project_id = $1
	`, projectID , time.Now())

	if err != nil { 
		return err 
	} 
	
	n, _ := res.RowsAffected() 
	log.Printf("Rows updated: %d", n) 
	return nil
}

func markFailed(ctx context.Context, db *sql.DB, projectID string, reason string) error {
	res, err := db.ExecContext(ctx, `
		UPDATE linked_projects
		SET status = 'ERROR',
		    error_reason = $2
		WHERE project_id = $1
	`, projectID, reason)

	if err != nil {
    log.Printf("Update failed: %v", err)
    return err
    }
    n, _ := res.RowsAffected()
    log.Printf("Rows updated: %d", n)


	return err
}