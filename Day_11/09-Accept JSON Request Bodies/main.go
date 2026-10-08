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

func createEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var employee Employee

	// Read JSON request body.
	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	employee.ID = 1

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Return created employee as JSON.
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

Header:
Content-Type: application/json

Body:

{
    "name": "Kiran",
    "age": 25,
    "email": "kiran@gmail.com"
}

Response:

{
    "id": 1,
    "name": "Kiran",
    "age": 25,
    "email": "kiran@gmail.com"
}

Important:

r.Body
    ↓
json.NewDecoder()
    ↓
Go struct

JSON → Go struct

If JSON is invalid:
400 Bad Request

If successful:
201 Created
*/