package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

var (
	mu        sync.RWMutex
	employees = map[int]Employee{
		1: {ID: 1, Name: "Vimal", Position: "Manager"},
	}
)

func updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	var updatedEmp Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Malformed JSON"})
		return
	}

	mu.Lock()
	_, exists := employees[id]
	if !exists {
		mu.Unlock()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Employee not found"})
		return
	}

	updatedEmp.ID = id
	employees[id] = updatedEmp
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEmp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /employee/{id}", updateEmployeeHandler)
	http.ListenAndServe(":8080", mux)
}