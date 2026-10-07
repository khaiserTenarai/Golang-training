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

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	employees = map[int]Employee{
		1: {ID: 1, Name: "John Doe", Department: "Engineering"},
		2: {ID: 2, Name: "Jane Smith", Department: "Marketing"},
	}
	mu sync.Mutex
)

func jsonError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	emp, exists := employees[id]
	mu.Unlock()

	if !exists {
		jsonError(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

func getEmployeeByDept(w http.ResponseWriter, r *http.Request) {
	dept := r.PathValue("dept")
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		jsonError(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	emp, exists := employees[id]
	mu.Unlock()

	if !exists || emp.Department != dept {
		jsonError(w, "Employee not found in department", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /employee/{id}", getEmployeeByID)
	mux.HandleFunc("GET /department/{dept}/employee/{id}", getEmployeeByDept)

	http.ListenAndServe(":8080", mux)
}