package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Dummy data for demonstration
	employees := []Employee{
		{ID: 1, Name: "Alice Smith", Email: "alice@example.com", Position: "Developer", Salary: 75000},
		{ID: 2, Name: "Bob Jones", Email: "bob@example.com", Position: "Manager", Salary: 95000},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

func main() {
	http.HandleFunc("/employee", getEmployees)
	http.ListenAndServe(":8080", nil)
}