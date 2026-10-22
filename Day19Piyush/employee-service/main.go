// Task 1: Employee Microservice | Task 3: CRUD | Task 4: Service-to-service | Task 5,6: DTOs
// Task 7: Externalized URLs | Task 8: Health | Task 9: Error handling | Task 10: Timeout
// Task 11: Retry | Task 12: Correlation ID | Task 13: Circuit breaker
// Run: go run ./employee-service   |   Port: 8081

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

	"go_microservices/config"
	"go_microservices/dto"
	"go_microservices/entity"
	"go_microservices/httpclient"
	"go_microservices/middleware"
)

var (
	mu         sync.RWMutex
	data       = map[int]entity.Employee{}
	next       = 1
	deptClient *httpclient.ServiceClient
)

func init() {
	for _, e := range []entity.Employee{
		{Name: "Piyush", Email: "piyush@test.com", Age: 25, DepartmentID: 1, Salary: 60000, IsActive: true},
		{Name: "Alice", Email: "alice@test.com", Age: 30, DepartmentID: 2, Salary: 55000, IsActive: true},
		{Name: "Bob", Email: "bob@test.com", Age: 28, DepartmentID: 3, Salary: 45000, IsActive: true},
		{Name: "Charlie", Email: "charlie@test.com", Age: 35, DepartmentID: 1, Salary: 80000, IsActive: false},
	} {
		e.ID = next
		e.CreatedAt = time.Now()
		data[e.ID] = e
		next++
	}
}

func toDTO(e entity.Employee) dto.EmployeeResponse {
	return dto.EmployeeResponse{
		ID: e.ID, Name: e.Name, Email: e.Email, Age: e.Age,
		DepartmentID: e.DepartmentID, Salary: e.Salary,
		IsActive: e.IsActive, CreatedAt: e.CreatedAt.Format(time.RFC3339),
	}
}

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func validate(req dto.CreateEmployeeRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return fmt.Errorf("name is required")
	}
	if !strings.Contains(req.Email, "@") {
		return fmt.Errorf("valid email required")
	}
	if req.Age < 18 || req.Age > 65 {
		return fmt.Errorf("age must be 18-65")
	}
	if req.Salary < 0 {
		return fmt.Errorf("salary cannot be negative")
	}
	return nil
}

func handleEmployees(w http.ResponseWriter, r *http.Request) {
	cid := middleware.GetCorrelationID(r.Context())
	switch r.Method {
	case "GET":
		mu.RLock()
		var list []dto.EmployeeResponse
		for _, e := range data {
			d := toDTO(e)
			d.Department = deptClient.GetDepartment(r.Context(), e.DepartmentID)
			list = append(list, d)
		}
		mu.RUnlock()
		respond(w, 200, dto.APIResponse{Success: true, Data: list, CorrelationID: cid})
	case "POST":
		var req dto.CreateEmployeeRequest
		json.NewDecoder(r.Body).Decode(&req)
		if err := validate(req); err != nil {
			respond(w, 400, dto.ErrorResponse{Error: err.Error(), Code: 400, CorrelationID: cid})
			return
		}
		mu.Lock()
		emp := entity.Employee{
			ID: next, Name: req.Name, Email: req.Email, Age: req.Age,
			DepartmentID: req.DepartmentID, Salary: req.Salary,
			IsActive: true, CreatedAt: time.Now(),
		}
		data[emp.ID] = emp
		next++
		mu.Unlock()
		d := toDTO(emp)
		d.Department = deptClient.GetDepartment(r.Context(), emp.DepartmentID)
		respond(w, 201, dto.APIResponse{Success: true, Data: d, CorrelationID: cid})
	}
}

func handleEmployeeByID(w http.ResponseWriter, r *http.Request) {
	cid := middleware.GetCorrelationID(r.Context())
	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])
	switch r.Method {
	case "GET":
		mu.RLock()
		emp, ok := data[id]
		mu.RUnlock()
		if !ok {
			respond(w, 404, dto.ErrorResponse{Error: "not found", Code: 404, CorrelationID: cid})
			return
		}
		d := toDTO(emp)
		d.Department = deptClient.GetDepartment(r.Context(), emp.DepartmentID)
		respond(w, 200, dto.APIResponse{Success: true, Data: d, CorrelationID: cid})
	case "PUT":
		var req dto.UpdateEmployeeRequest
		json.NewDecoder(r.Body).Decode(&req)
		mu.Lock()
		old, ok := data[id]
		if !ok {
			mu.Unlock()
			respond(w, 404, dto.ErrorResponse{Error: "not found", Code: 404, CorrelationID: cid})
			return
		}
		emp := entity.Employee{
			ID: id, Name: req.Name, Email: req.Email, Age: req.Age,
			DepartmentID: req.DepartmentID, Salary: req.Salary,
			IsActive: req.IsActive, CreatedAt: old.CreatedAt,
		}
		data[id] = emp
		mu.Unlock()
		respond(w, 200, dto.APIResponse{Success: true, Data: toDTO(emp), CorrelationID: cid})
	case "DELETE":
		mu.Lock()
		if _, ok := data[id]; !ok {
			mu.Unlock()
			respond(w, 404, dto.ErrorResponse{Error: "not found", Code: 404, CorrelationID: cid})
			return
		}
		delete(data, id)
		mu.Unlock()
		respond(w, 200, dto.APIResponse{Success: true, Message: "deleted", CorrelationID: cid})
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"service": "employee-service", "status": "healthy", "version": "1.0.0",
		"dependencies": map[string]string{"department-service": "circuit-breaker: " + deptClient.CircuitState()},
	})
}

func main() {
	cfg := config.Load()
	deptClient = httpclient.NewServiceClient(cfg.DepartmentServiceURL, cfg.HTTPTimeoutSecs)
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/employees", handleEmployees)
	mux.HandleFunc("/employees/", handleEmployeeByID)
	handler := middleware.Logging(middleware.CorrelationID(mux))
	fmt.Println("[Employee-Service] Running on http://localhost:8081")
	fmt.Printf("[Employee-Service] Department URL: %s\n", cfg.DepartmentServiceURL)
	log.Fatal(http.ListenAndServe(":8081", handler))
}
