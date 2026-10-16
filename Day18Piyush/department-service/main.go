// Task 5: Define Department Service
// Task 7: Database-per-service (own in-memory store)
// Run: go run main.go
// Port: 8082

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type Department struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

var (
	mu   sync.RWMutex
	data = map[int]Department{
		1: {1, "Engineering", "Building A"},
		2: {2, "Design", "Building B"},
		3: {3, "Marketing", "Building C"},
		4: {4, "HR", "Building D"},
		5: {5, "Finance", "Building A"},
	}
	next = 6
)

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func handleDepartments(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mu.RLock()
		var list []Department
		for _, d := range data {
			list = append(list, d)
		}
		mu.RUnlock()
		respond(w, 200, map[string]interface{}{"service": "department-service", "data": list})

	case "POST":
		var dept Department
		json.NewDecoder(r.Body).Decode(&dept)
		mu.Lock()
		dept.ID = next
		next++
		data[dept.ID] = dept
		mu.Unlock()
		log.Printf("[EVENT] department.created: ID=%d Name=%s", dept.ID, dept.Name)
		respond(w, 201, map[string]interface{}{"data": dept})
	}
}

func handleDepartmentByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])
	mu.RLock()
	dept, ok := data[id]
	mu.RUnlock()
	if !ok {
		respond(w, 404, map[string]interface{}{"error": "department not found"})
		return
	}
	respond(w, 200, map[string]interface{}{"data": dept})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, 200, map[string]interface{}{"service": "department-service", "status": "healthy", "port": 8082})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/departments", handleDepartments)
	mux.HandleFunc("/departments/", handleDepartmentByID)

	fmt.Println("[Department-Service] Running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
