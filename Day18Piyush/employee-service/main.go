// Task 4: Define Employee Service
// Task 7: Database-per-service (own in-memory store)
// Task 8: Synchronous REST communication (calls department-service)
// Run: go run main.go
// Port: 8081

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Employee struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	DepartmentID int       `json:"department_id"`
	Salary       float64   `json:"salary"`
	CreatedAt    time.Time `json:"created_at"`
}

type EmployeeWithDept struct {
	Employee
	Department interface{} `json:"department,omitempty"`
}

var (
	mu   sync.RWMutex
	data = map[int]Employee{}
	next = 1
)

func init() {
	seed := []Employee{
		{Name: "Piyush", Email: "piyush@test.com", Age: 25, DepartmentID: 1, Salary: 60000},
		{Name: "Alice", Email: "alice@test.com", Age: 30, DepartmentID: 2, Salary: 55000},
		{Name: "Bob", Email: "bob@test.com", Age: 28, DepartmentID: 1, Salary: 45000},
		{Name: "Charlie", Email: "charlie@test.com", Age: 35, DepartmentID: 3, Salary: 80000},
	}
	for _, e := range seed {
		e.ID = next
		e.CreatedAt = time.Now()
		data[e.ID] = e
		next++
	}
}

// Task 8: Synchronous REST call to Department Service
func fetchDepartment(deptID int) interface{} {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8082/departments/%d", deptID))
	if err != nil {
		log.Printf("[Employee-Service] Failed to reach department-service: %v", err)
		return nil
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	if d, ok := result["data"]; ok {
		return d
	}
	return nil
}

func respond(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func handleEmployees(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mu.RLock()
		var list []EmployeeWithDept
		for _, e := range data {
			ed := EmployeeWithDept{Employee: e, Department: fetchDepartment(e.DepartmentID)}
			list = append(list, ed)
		}
		mu.RUnlock()
		respond(w, 200, map[string]interface{}{"service": "employee-service", "data": list, "count": len(list)})

	case "POST":
		var emp Employee
		json.NewDecoder(r.Body).Decode(&emp)
		if emp.Name == "" || emp.Email == "" {
			respond(w, 400, map[string]interface{}{"error": "name and email required"})
			return
		}
		mu.Lock()
		emp.ID = next
		next++
		emp.CreatedAt = time.Now()
		data[emp.ID] = emp
		mu.Unlock()

		// Task 9: Publish event (simulated async)
		log.Printf("[EVENT] employee.created: ID=%d Name=%s", emp.ID, emp.Name)
		respond(w, 201, map[string]interface{}{"data": emp})
	}
}

func handleEmployeeByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	id, _ := strconv.Atoi(parts[len(parts)-1])

	switch r.Method {
	case "GET":
		mu.RLock()
		emp, ok := data[id]
		mu.RUnlock()
		if !ok {
			respond(w, 404, map[string]interface{}{"error": "not found"})
			return
		}
		ed := EmployeeWithDept{Employee: emp, Department: fetchDepartment(emp.DepartmentID)}
		respond(w, 200, map[string]interface{}{"data": ed})

	case "DELETE":
		mu.Lock()
		if _, ok := data[id]; !ok {
			mu.Unlock()
			respond(w, 404, map[string]interface{}{"error": "not found"})
			return
		}
		delete(data, id)
		mu.Unlock()
		log.Printf("[EVENT] employee.deleted: ID=%d", id)
		respond(w, 200, map[string]interface{}{"message": "deleted"})
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	respond(w, 200, map[string]interface{}{"service": "employee-service", "status": "healthy", "port": 8081})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/employees", handleEmployees)
	mux.HandleFunc("/employees/", handleEmployeeByID)

	fmt.Println("[Employee-Service] Running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
