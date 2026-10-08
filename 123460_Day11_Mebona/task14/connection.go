package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "postgres://postgres:pgadmin@localhost:5432/employee_db?sslmode=disable"
	
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	if err := db.Ping(); err != nil {
		http.Error(w, "Database is down", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Database connection is healthy"))
}

func main() {
	initDB()
	defer db.Close()

	http.HandleFunc("/health", healthCheckHandler)
	fmt.Println("Server running on :8080...")
	http.ListenAndServe(":8080", nil)
}