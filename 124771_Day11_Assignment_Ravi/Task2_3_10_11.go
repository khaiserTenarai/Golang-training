package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

// Temporary in-memory data
var employees = []Employee{
	{ID: 1, Name: "Ravi", Email: "ravi@example.com", Department: "IT"},
	{ID: 2, Name: "Amit", Email: "amit@example.com", Department: "HR"},
	{ID: 3, Name: "Priya", Email: "priya@example.com", Department: "IT"},
}

// QUESTION 2 + QUESTION 11
// GET /employees
// Q2: Return all employees.
// Q11: Support query parameters, such as ?department=IT.
func getEmployees(w http.ResponseWriter, r *http.Request) {

	// Q11: Read the department query parameter.
	// Example: /employees?department=IT
	department := r.URL.Query().Get("department")
	result := employees
	if department != "" {
		result = []Employee{}

		for _, employee := range employees {
			if strings.EqualFold(employee.Department, department) {
				result = append(result, employee)
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// QUESTION 3 + QUESTION 10
// GET /employees/{id}
// Q3: Return details of a particular employee.
// Q10: Extract the employee ID from the URL path.
func getEmployeeByID(w http.ResponseWriter, r *http.Request) {
	// Q10: Extract the path parameter.
	idString := strings.TrimPrefix(r.URL.Path, "/employees/")
	id, err := strconv.Atoi(idString)
	if err != nil || id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid employee ID",
		})
		return
	}

	// Q3: Search for the employee using the ID.
	for _, employee := range employees {
		if employee.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(employee)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "Employee not found",
	})
}

func main() {
	// QUESTION 2
	// Register the GET /employees endpoint.
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		getEmployees(w, r)
	})

	// QUESTION 3 + QUESTION 10
	// Register the GET /employees/{id} endpoint.
	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		getEmployeeByID(w, r)
	})

	// QUESTION 1: Start the HTTP server (from Task 1).
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
