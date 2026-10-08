package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
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

// PUT /employee/{id}
func updateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get ID from URL
	idString := strings.TrimPrefix(r.URL.Path, "/employee/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var updatedEmployee Employee

	// Read JSON body
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range employees {

		if employees[i].ID == id {

			updatedEmployee.ID = id
			employees[i] = updatedEmployee

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedEmployee)

			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/employee/", updateEmployee)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
Test:

PUT http://localhost:8080/employee/1

Body:

{
    "name": "Ganesh Reddy",
    "age": 23,
    "email": "ganeshreddy@gmail.com"
}

Response:

{
    "id": 1,
    "name": "Ganesh Reddy",
    "age": 23,
    "email": "ganeshreddy@gmail.com"
}

Important:
strconv.Atoi() -> converts string ID to int
json.NewDecoder() -> reads request JSON
PUT -> updates an existing resource
404 -> employee not found
*/