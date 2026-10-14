package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

// In-memory employee storage
var employees []Employee
var nextID = 1

// Q12: Create logging middleware
// Logs the HTTP method, path, and time taken by each request.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		log.Printf("Started: %s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		log.Printf(
			"Completed: %s %s | Duration: %v",
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}

// Q13: Validate employee data
func validateEmployee(e Employee) error {
	if strings.TrimSpace(e.Name) == "" {
		return fmt.Errorf("name is required")
	}

	if strings.TrimSpace(e.Email) == "" {
		return fmt.Errorf("email is required")
	}

	if !strings.Contains(e.Email, "@") {
		return fmt.Errorf("email must contain @")
	}

	if strings.TrimSpace(e.Department) == "" {
		return fmt.Errorf("department is required")
	}

	return nil
}

// Helper function to return JSON responses
func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// GET /employees
// Lists all employees
func getEmployees(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "Method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, employees)
}

// POST /employees
// Creates an employee after validating the request data
func createEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "Method not allowed",
		})
		return
	}

	var employee Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "Invalid JSON request body",
		})
		return
	}

	// Q13: Validate employee data before saving
	if err := validateEmployee(employee); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
		return
	}

	employee.ID = nextID
	nextID++

	employees = append(employees, employee)

	writeJSON(w, http.StatusCreated, employee)
}

func main() {
	mux := http.NewServeMux()

	// Register endpoints
	mux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getEmployees(w, r)
		case http.MethodPost:
			createEmployee(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{
				"error": "Method not allowed",
			})
		}
	})

	// Q12: Apply logging middleware to all requests
	loggedMux := loggingMiddleware(mux)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", loggedMux))
}
