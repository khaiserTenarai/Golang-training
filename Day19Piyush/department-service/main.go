// Task 2: Department Microservice | Task 3: CRUD | Task 5,6: DTOs | Task 8: Health | Task 12: Correlation ID
// Run: go run ./department-service   |   Port: 8082

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"go_microservices/dto"
	"go_microservices/entity"
	"go_microservices/middleware"
)

var (
	mu   sync.RWMutex
	data = map[int]entity.Department{}
	next = 1
)

func init() {
	for _, d := range []entity.Department{
		{Name: "Engineering", Location: "Building A", Budget: 500000},
		{Name: "Design", Location: "Building B", Budget: 300000},
		{Name: "Marketing", Location: "Building C", Budget: 250000},
		{Name: "HR", Location: "Building D", Budget: 200000},
	} {
		d.ID = next
		data[d.ID] = d
		next++
	}
}

func toDTO(d entity.Department) dto.DepartmentResponse {
	return dto.DepartmentResponse{ID: d.ID, Name: d.Name, Location: d.Location, Budget: d.Budget}
}

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func handleDepartments(w http.ResponseWriter, r *http.Request) {
	cid := middleware.GetCorrelationID(r.Context())
	switch r.Method {
	case "GET":
		mu.RLock()
		var list []dto.DepartmentResponse
		for _, d := range data {
			list = append(list, toDTO(d))
		}
		mu.RUnlock()
		respond(w, 200, dto.APIResponse{Success: true, Data: list, CorrelationID: cid})
	case "POST":
		var req dto.CreateDepartmentRequest
		json.NewDecoder(r.Body).Decode(&req)
		if req.Name == "" {
			respond(w, 400, dto.ErrorResponse{Error: "name required", Code: 400, CorrelationID: cid})
			return
		}
		mu.Lock()
		dept := entity.Department{ID: next, Name: req.Name, Location: req.Location, Budget: req.Budget}
		data[dept.ID] = dept
		next++
		mu.Unlock()
		respond(w, 201, dto.APIResponse{Success: true, Data: toDTO(dept), CorrelationID: cid})
	}
}

func handleDepartmentByID(w http.ResponseWriter, r *http.Request) {
	cid := middleware.GetCorrelationID(r.Context())
	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])
	switch r.Method {
	case "GET":
		mu.RLock()
		dept, ok := data[id]
		mu.RUnlock()
		if !ok {
			respond(w, 404, dto.ErrorResponse{Error: "not found", Code: 404, CorrelationID: cid})
			return
		}
		respond(w, 200, dto.APIResponse{Success: true, Data: toDTO(dept), CorrelationID: cid})
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
	json.NewEncoder(w).Encode(dto.HealthResponse{Service: "department-service", Status: "healthy", Version: "1.0.0"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/departments", handleDepartments)
	mux.HandleFunc("/departments/", handleDepartmentByID)
	handler := middleware.Logging(middleware.CorrelationID(mux))
	fmt.Println("[Department-Service] Running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", handler))
}
