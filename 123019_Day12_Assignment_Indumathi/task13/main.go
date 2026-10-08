package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{
		Status:    "UP",
		Timestamp: time.Now(),
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)
	http.ListenAndServe(":8080", mux)
}