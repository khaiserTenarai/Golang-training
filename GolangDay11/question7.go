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
	employees = make(map[int]Employee)
	mu        sync.Mutex
	nextID    = 1
)

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
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

func getEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	mu.Lock()
	emp, exists := employees[id]
	mu.Unlock()

	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	mu.Lock()
	if _, exists := employees[id]; !exists {
		mu.Unlock()
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	emp.ID = id
	employees[id] = emp
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	mu.Lock()
	if _, exists := employees[id]; !exists {
		mu.Unlock()
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	delete(employees, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /employee", createEmployee)
	mux.HandleFunc("GET /employee/{id}", getEmployee)
	mux.HandleFunc("PUT /employee/{id}", updateEmployee)
	mux.HandleFunc("DELETE /employee/{id}", deleteEmployee)

	http.ListenAndServe(":8080", mux)
}