package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

var (
	mu        sync.Mutex
	employees = make(map[int]Employee)
	nextID    = 1
)

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON payload"})
		return
	}

	mu.Lock()
	emp.ID = nextID
	nextID++
	employees[emp.ID] = emp
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(emp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /employee", createEmployeeHandler)
	http.ListenAndServe(":8080", mux)
}