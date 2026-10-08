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
	2: {ID: 2, Name: "Bob Jones", Email: "bob@example.com", Position: "Manager", Salary: 95000},
}

func employeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	// Expecting path like /employee/1
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

	employee, exists := mockDB[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employee)
}

func main() {
	http.HandleFunc("/employee/", employeeDetailHandler)
	http.ListenAndServe(":8080", nil)
}