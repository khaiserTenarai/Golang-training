package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "host=localhost port=5432 user=postgres password=secret dbname=company sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Database connection configuration error: %v", err)
	}
	defer db.Close()

	// Verify database connection alive status
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot reach PostgreSQL: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL database!")
}