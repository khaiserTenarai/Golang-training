package main

import (
	"api/handler"
	"api/middleware"
	"api/repository"
	"api/service"
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq" 
)

func main() {
	
	connStr := "user=postgres password=pgadmin dbname='Day 12' sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
	}
	defer db.Close()

	
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	log.Println("Successfully connected to PostgreSQL!")

	
	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	empHandler := handler.NewEmployeeHandler(svc)

	mux := http.NewServeMux()

	
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("OK")) })
	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("READY")) })
	
	
	mux.Handle("/api/employees", middleware.AuthMiddleware(empHandler))

	
	loggedMux := middleware.RequestIDAndLoggingMiddleware(mux)

	log.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", loggedMux)
}