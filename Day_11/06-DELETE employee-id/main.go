package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID    int
	Name  string
	Age   int
	Email string
}

var employees = []Employee{
	{ID: 1, Name: "Ganesh", Age: 22, Email: "ganesh@gmail.com"},
	{ID: 2, Name: "Ravi", Age: 23, Email: "ravi@gmail.com"},
}

// DELETE /employee/{id}
func deleteEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
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

	for i, employee := range employees {

		if employee.ID == id {

			// Remove employee from slice
			employees = append(employees[:i], employees[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/employee/", deleteEmployee)

	fmt.Println("Server started on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

/*
Test:

DELETE http://localhost:8080/employee/1

If employee exists:
Status -> 204 No Content

If employee doesn't exist:

DELETE http://localhost:8080/employee/10

Status -> 404 Not Found

Important:
DELETE -> removes a resource
204 -> successful deletion with no response body
404 -> employee not found
*/