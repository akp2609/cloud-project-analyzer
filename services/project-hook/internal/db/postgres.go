package db

import (
	"log"
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error){
	dsn := os.Getenv("DATABASE_URL")
	log.Printf("DATABASE_URL=%s", dsn)
	return sql.Open("postgres", dsn)
}