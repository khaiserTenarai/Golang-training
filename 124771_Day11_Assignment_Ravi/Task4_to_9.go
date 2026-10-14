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

var employees = []Employee{
	{1, "Ravi", "ravi@example.com", "IT"},
	{2, "Amit", "amit@example.com", "HR"},
}

// Q2 + Q7+ Q8: GET all employees and return JSON
func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // Q7: 200 OK
	//Q8 — Return JSON responses
	json.NewEncoder(w).Encode(employees)
}

// Q4 + Q9: POST - create employee from JSON
func createEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee

	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	e.ID = len(employees) + 1
	employees = append(employees, e)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

// Q5: PUT - update employee by ID
func updateEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employees/"))

	var updated Employee
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, e := range employees {
		if e.ID == id {
			updated.ID = id
			employees[i] = updated

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

// Q6: DELETE - remove employee by ID
func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employees/"))

	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	// Basic HTTP server
	// GET and POST use /employees
	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getEmployees(w, r)
		case http.MethodPost:
			createEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// PUT and DELETE use /employees/{id}
	http.HandleFunc("/employees/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			updateEmployee(w, r)
		case http.MethodDelete:
			deleteEmployee(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
