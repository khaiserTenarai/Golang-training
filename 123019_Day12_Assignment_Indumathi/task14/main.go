package main

import (
	"encoding/json"
	"net/http"
)

type Database struct {
	IsConnected bool
}

var db = &Database{IsConnected: true}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Check downstream infrastructure connectivity
	if !db.IsConnected {
		w.WriteHeader(http.StatusServiceUnavailable)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "DOWN",
			"reason": "Database connection unavailable",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "READY",
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ready", readinessHandler)
	http.ListenAndServe(":8080", mux)
}