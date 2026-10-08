package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var employees = []Employee{
	{ID: 1, Name: "Ganesh", Age: 22, Email: "ganesh@gmail.com"},
	{ID: 2, Name: "Ravi", Age: 23, Email: "ravi@gmail.com"},
}

// POST /employee
func createEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var employee Employee

	// Decode JSON request body
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Generate ID automatically
	employee.ID = len(employees) + 1

	employees = append(employees, employee)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send created employee as JSON
	json.NewEncoder(w).Encode(employee)
}

func main() {

	http.HandleFunc("/employee", createEmployee)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
Test:

POST http://localhost:8080/employee

Body:
{
    "name": "Kiran",
    "age": 25,
    "email": "kiran@gmail.com"
}

Response:
{
    "id": 3,
    "name": "Kiran",
    "age": 25,
    "email": "kiran@gmail.com"
}

Important:
json.NewDecoder() -> JSON to Go struct
json.NewEncoder() -> Go struct to JSON
201 Created -> successful creation
*/