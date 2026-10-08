package main

import (
	"log"
	"net/http"
)

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	// Use store instead of db
	store := NewPostgresStore(db)
	server := &Server{store: store}

	mux := http.NewServeMux()
	mux.HandleFunc("/employees", server.handleEmployees)
	mux.HandleFunc("/employees/{id}", server.handleEmployeeByID)

	loggedRouter := loggingMiddleware(mux)

	log.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", loggedRouter); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}