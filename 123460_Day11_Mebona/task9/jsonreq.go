package main

import (
	"encoding/json"
	"net/http"
)

type Payload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func acceptJSONHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var p Payload
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&p); err != nil {
		http.Error(w, "Bad request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func main() {
	http.HandleFunc("/submit", acceptJSONHandler)
	http.ListenAndServe(":8080", nil)
}