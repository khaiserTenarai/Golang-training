// Task 1: Create a Basic HTTP Server
// Run: go run main.go
// Test: open http://localhost:8080 in browser
//       or use: curl http://localhost:8080

package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Day 11 — REST API with Go!")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status": "healthy", "message": "Server is running"}`)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Day 11 REST API Training — Piyush")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Task 1: Basic HTTP Server")
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("Endpoints:")
	fmt.Println("  GET /        → Home page")
	fmt.Println("  GET /health  → Health check (JSON)")
	fmt.Println("  GET /about   → About page")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
