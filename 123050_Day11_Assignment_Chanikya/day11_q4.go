package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var employees []Employee
var nextID = 1

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	employee.ID = nextID
	nextID++

	employees = append(employees, employee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(employee)
}

func main() {
	http.HandleFunc("/employee", createEmployee)

	http.ListenAndServe(":8080", nil)
}
