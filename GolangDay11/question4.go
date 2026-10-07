package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var (
	employees = make(map[int]Employee)
	mu        sync.Mutex
	nextID    = 1
)

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	emp.ID = nextID
	employees[nextID] = emp
	nextID++
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	mux := http.NewServeMux()
	
	mux.HandleFunc("POST /employee", createEmployeeHandler)
	
	http.ListenAndServe(":8080", mux)
}