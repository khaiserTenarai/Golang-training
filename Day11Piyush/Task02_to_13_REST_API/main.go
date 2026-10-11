// Tasks 2–13: Complete In-Memory REST API (no PostgreSQL)
// Task 2  — GET /employees
// Task 3  — GET /employees/{id} (employee detail)
// Task 4  — POST /employees
// Task 5  — PUT /employees/{id}
// Task 6  — DELETE /employees/{id}
// Task 7  — Correct HTTP status codes
// Task 8  — JSON responses
// Task 9  — Accept JSON request bodies
// Task 10 — Path parameters
// Task 11 — Query parameters (?name=, ?department=, ?min_salary=, ?page=, &page_size=)
// Task 12 — Logging middleware
// Task 13 — Validate employee data
//
// Run: go run main.go
// Test with curl (see bottom of this file for examples)

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// ============================================================
// Models (Task 8: JSON tags for proper responses)
// ============================================================

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ============================================================
// In-Memory Store (thread-safe)
// ============================================================

type EmployeeStore struct {
	mu        sync.RWMutex
	employees map[int]Employee
	nextID    int
}

func NewStore() *EmployeeStore {
	return &EmployeeStore{
		employees: map[int]Employee{
			1: {1, "Piyush", "piyush@test.com", 25, "Engineering", 60000},
			2: {2, "Alice", "alice@test.com", 30, "Design", 55000},
			3: {3, "Bob", "bob@test.com", 28, "Marketing", 45000},
		},
		nextID: 4,
	}
}

// ============================================================
// Task 13: Validate Employee Data
// ============================================================

func validateEmployee(emp Employee) []ValidationError {
	var errs []ValidationError

	if strings.TrimSpace(emp.Name) == "" {
		errs = append(errs, ValidationError{"name", "Name is required"})
	} else if len(emp.Name) < 2 {
		errs = append(errs, ValidationError{"name", "Name must be at least 2 characters"})
	}

	if strings.TrimSpace(emp.Email) == "" {
		errs = append(errs, ValidationError{"email", "Email is required"})
	} else if !strings.Contains(emp.Email, "@") || !strings.Contains(emp.Email, ".") {
		errs = append(errs, ValidationError{"email", "Invalid email format"})
	}

	if emp.Age < 18 || emp.Age > 65 {
		errs = append(errs, ValidationError{"age", "Age must be between 18 and 65"})
	}

	if strings.TrimSpace(emp.Department) == "" {
		errs = append(errs, ValidationError{"department", "Department is required"})
	}

	if emp.Salary < 0 {
		errs = append(errs, ValidationError{"salary", "Salary cannot be negative"})
	}

	return errs
}

// ============================================================
// Task 12: Logging Middleware
// ============================================================

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := &statusWriter{ResponseWriter: w, statusCode: 200}
		next.ServeHTTP(lw, r)
		log.Printf("[%s] %s %s → %d (%v)",
			r.Method, r.URL.Path, r.RemoteAddr, lw.statusCode, time.Since(start).Round(time.Microsecond))
	})
}

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// ============================================================
// Helper: Send JSON Response (Tasks 7, 8)
// ============================================================

func sendJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // Task 7: correct status code
	json.NewEncoder(w).Encode(payload)
}

// ============================================================
// Task 10: Extract Path Parameter (manual, no external router)
// ============================================================

func extractID(path, prefix string) (int, error) {
	idStr := strings.TrimPrefix(path, prefix)
	idStr = strings.TrimSuffix(idStr, "/")
	return strconv.Atoi(idStr)
}

// ============================================================
// Handlers
// ============================================================

func (s *EmployeeStore) handleEmployees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	// Task 2: GET /employees (list all)
	// Task 11: Query parameters for filtering, search, pagination
	case http.MethodGet:
		s.mu.RLock()
		defer s.mu.RUnlock()

		query := r.URL.Query()
		nameFilter := strings.ToLower(query.Get("name"))
		deptFilter := strings.ToLower(query.Get("department"))
		minSalary, _ := strconv.ParseFloat(query.Get("min_salary"), 64)
		maxSalary, _ := strconv.ParseFloat(query.Get("max_salary"), 64)
		page, _ := strconv.Atoi(query.Get("page"))
		pageSize, _ := strconv.Atoi(query.Get("page_size"))

		if page < 1 {
			page = 1
		}
		if pageSize < 1 {
			pageSize = 10
		}

		var filtered []Employee
		for _, emp := range s.employees {
			if nameFilter != "" && !strings.Contains(strings.ToLower(emp.Name), nameFilter) {
				continue
			}
			if deptFilter != "" && !strings.EqualFold(emp.Department, deptFilter) {
				continue
			}
			if minSalary > 0 && emp.Salary < minSalary {
				continue
			}
			if maxSalary > 0 && emp.Salary > maxSalary {
				continue
			}
			filtered = append(filtered, emp)
		}

		// Pagination
		total := len(filtered)
		start := (page - 1) * pageSize
		end := start + pageSize
		if start > total {
			start = total
		}
		if end > total {
			end = total
		}
		paged := filtered[start:end]

		sendJSON(w, http.StatusOK, map[string]interface{}{
			"success":    true,
			"data":       paged,
			"total":      total,
			"page":       page,
			"page_size":  pageSize,
		})

	// Task 4: POST /employees (Task 9: accept JSON body)
	case http.MethodPost:
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			sendJSON(w, http.StatusBadRequest, APIResponse{false, "Invalid JSON body", nil})
			return
		}

		// Task 13: Validate
		if errs := validateEmployee(emp); len(errs) > 0 {
			sendJSON(w, http.StatusUnprocessableEntity, APIResponse{false, "Validation failed", errs})
			return
		}

		s.mu.Lock()
		emp.ID = s.nextID
		s.nextID++
		s.employees[emp.ID] = emp
		s.mu.Unlock()

		sendJSON(w, http.StatusCreated, APIResponse{true, "Employee created", emp}) // 201

	default:
		sendJSON(w, http.StatusMethodNotAllowed, APIResponse{false, "Method not allowed", nil}) // 405
	}
}

func (s *EmployeeStore) handleEmployeeByID(w http.ResponseWriter, r *http.Request) {
	// Task 10: Extract path parameter
	id, err := extractID(r.URL.Path, "/employees/")
	if err != nil {
		sendJSON(w, http.StatusBadRequest, APIResponse{false, "Invalid employee ID", nil}) // 400
		return
	}

	switch r.Method {

	// Task 3: GET /employees/{id} (detail endpoint)
	case http.MethodGet:
		s.mu.RLock()
		emp, exists := s.employees[id]
		s.mu.RUnlock()
		if !exists {
			sendJSON(w, http.StatusNotFound, APIResponse{false, "Employee not found", nil}) // 404
			return
		}
		sendJSON(w, http.StatusOK, APIResponse{true, "Employee found", emp}) // 200

	// Task 5: PUT /employees/{id}
	case http.MethodPut:
		var emp Employee
		if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
			sendJSON(w, http.StatusBadRequest, APIResponse{false, "Invalid JSON body", nil})
			return
		}

		// Task 13: Validate
		if errs := validateEmployee(emp); len(errs) > 0 {
			sendJSON(w, http.StatusUnprocessableEntity, APIResponse{false, "Validation failed", errs})
			return
		}

		s.mu.Lock()
		if _, exists := s.employees[id]; !exists {
			s.mu.Unlock()
			sendJSON(w, http.StatusNotFound, APIResponse{false, "Employee not found", nil})
			return
		}
		emp.ID = id
		s.employees[id] = emp
		s.mu.Unlock()

		sendJSON(w, http.StatusOK, APIResponse{true, "Employee updated", emp})

	// Task 6: DELETE /employees/{id}
	case http.MethodDelete:
		s.mu.Lock()
		if _, exists := s.employees[id]; !exists {
			s.mu.Unlock()
			sendJSON(w, http.StatusNotFound, APIResponse{false, "Employee not found", nil})
			return
		}
		delete(s.employees, id)
		s.mu.Unlock()

		sendJSON(w, http.StatusOK, APIResponse{true, "Employee deleted", nil})

	default:
		sendJSON(w, http.StatusMethodNotAllowed, APIResponse{false, "Method not allowed", nil})
	}
}

// ============================================================
// Router
// ============================================================

func (s *EmployeeStore) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/employees" || r.URL.Path == "/employees/" {
		s.handleEmployees(w, r)
	} else if strings.HasPrefix(r.URL.Path, "/employees/") {
		s.handleEmployeeByID(w, r)
	} else if r.URL.Path == "/" {
		sendJSON(w, http.StatusOK, APIResponse{true, "Employee REST API — Day 11", nil})
	} else {
		sendJSON(w, http.StatusNotFound, APIResponse{false, "Route not found", nil})
	}
}

// ============================================================
// Main
// ============================================================

func main() {
	store := NewStore()

	// Task 12: Wrap with logging middleware
	handler := loggingMiddleware(store)

	fmt.Println("Tasks 2-13: In-Memory Employee REST API")
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println()
	fmt.Println("Endpoints:")
	fmt.Println("  GET    /employees                          → List all (Task 2)")
	fmt.Println("  GET    /employees?name=piyush              → Search by name (Task 11)")
	fmt.Println("  GET    /employees?department=Engineering    → Filter by dept (Task 11)")
	fmt.Println("  GET    /employees?min_salary=50000         → Filter by salary (Task 11)")
	fmt.Println("  GET    /employees?page=1&page_size=2       → Pagination (Task 11)")
	fmt.Println("  GET    /employees/{id}                     → Detail (Tasks 3, 10)")
	fmt.Println("  POST   /employees                          → Create (Tasks 4, 9)")
	fmt.Println("  PUT    /employees/{id}                     → Update (Task 5)")
	fmt.Println("  DELETE /employees/{id}                     → Delete (Task 6)")
	fmt.Println()
	fmt.Println("Features: Status codes (7), JSON (8), Body parsing (9),")
	fmt.Println("  Path params (10), Query params (11), Logging (12), Validation (13)")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

/*
============================================================
  CURL TEST COMMANDS
============================================================

# List all employees (Task 2)
curl http://localhost:8080/employees

# Search by name (Task 11)
curl "http://localhost:8080/employees?name=piyush"

# Filter by department (Task 11)
curl "http://localhost:8080/employees?department=Engineering"

# Filter by salary range (Task 11)
curl "http://localhost:8080/employees?min_salary=50000&max_salary=70000"

# Pagination (Task 11)
curl "http://localhost:8080/employees?page=1&page_size=2"

# Get single employee (Tasks 3, 10)
curl http://localhost:8080/employees/1

# Create employee (Tasks 4, 9, 13)
curl -X POST http://localhost:8080/employees -H "Content-Type: application/json" -d "{\"name\":\"Diana\",\"email\":\"diana@test.com\",\"age\":26,\"department\":\"HR\",\"salary\":52000}"

# Create with validation error (Task 13)
curl -X POST http://localhost:8080/employees -H "Content-Type: application/json" -d "{\"name\":\"\",\"email\":\"bad\",\"age\":10,\"department\":\"\",\"salary\":-100}"

# Update employee (Task 5)
curl -X PUT http://localhost:8080/employees/1 -H "Content-Type: application/json" -d "{\"name\":\"Piyush Updated\",\"email\":\"piyush@new.com\",\"age\":26,\"department\":\"Backend\",\"salary\":70000}"

# Delete employee (Task 6)
curl -X DELETE http://localhost:8080/employees/3

# Not found (Task 7 - 404)
curl http://localhost:8080/employees/999
*/
