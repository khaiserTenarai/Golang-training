package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

var mockDB = map[int]Employee{
	1: {ID: 1, Name: "Alice Smith", Email: "alice@example.com", Position: "Developer", Salary: 75000},
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	if _, exists := mockDB[id]; !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	var updatedEmp Employee
	if err := json.NewDecoder(r.Body).Decode(&updatedEmp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedEmp.ID = id
	mockDB[id] = updatedEmp

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedEmp)
}

func main() {
	http.HandleFunc("/employee/", updateEmployee)
	http.ListenAndServe(":8080", nil)
}