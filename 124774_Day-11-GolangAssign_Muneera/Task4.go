package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee3 struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var employees = []Employee3{
	{ID: 1, Name: "Muneera", Department: "IT", Salary: 50000},
}

func createEmployee(w http.ResponseWriter, r *http.Request) {

	var employee Employee3

	// Get JSON from request body
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Add employee to the slice
	employees = append(employees, employee)

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(employee)
}

func main() {

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Server running on http://localhost:6060")

	err := http.ListenAndServe(":6060", nil)

	if err != nil {
		fmt.Println(err)
	}
}

// Request URL:
// http://localhost:8080/employee

// Body:
// {
//   "id": 4,
//   "name": "Arun",
//   "department": "IT",
//   "salary": 55000
// }
