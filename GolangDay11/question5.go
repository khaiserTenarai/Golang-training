package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var (
	employees = map[int]Employee{
		1: {ID: 1, Name: "Lakshmi Shibu", Department: "Engineering"},
	}
	mu sync.Mutex
)

func updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var updatedEmp Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmp); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	if _, exists := employees[id]; !exists {
		mu.Unlock()
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	updatedEmp.ID = id
	employees[id] = updatedEmp
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEmp)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("PUT /employee/{id}", updateEmployeeHandler)

	http.ListenAndServe(":8080", mux)
}