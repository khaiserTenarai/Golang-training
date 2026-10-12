package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee13 struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func createEmployee(w http.ResponseWriter, r *http.Request) {

	var employee Employee13

	// Read JSON
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate employee data
	if employee.ID <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if employee.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	if employee.Department == "" {
		http.Error(w, "Department is required", http.StatusBadRequest)
		return
	}

	if employee.Salary <= 0 {
		http.Error(w, "Salary must be greater than 0", http.StatusBadRequest)
		return
	}

	// Data is valid
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func main() {

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Server running on http://localhost:3030")

	http.ListenAndServe(":3030", nil)
}

//http://localhost:3030/employee

//JSON Body:
// {
//   "id": 1,
//   "name": "Muneera",
//   "department": "IT",
//   "salary": 50000
// }

// 2. Invalid ID
// {
//   "id": 0,
//   "name": "Muneera",
//   "department": "IT",
//   "salary": 50000
// }
