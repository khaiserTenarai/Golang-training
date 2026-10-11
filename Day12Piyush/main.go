// Day 12 — Production REST API Architecture
// Task 15: Mini Project — All 15 tasks combined into one production-style layered API.
//
// Architecture (Tasks 1-4):
//   main.go → routes → middleware → handler → service → repository
//   Each layer depends on interfaces (dependency injection).
//
// Run: go mod tidy && go run main.go
//
// No PostgreSQL needed — uses in-memory store.

package main

import (
	"fmt"
	"log"
	"net/http"

	"production_api/handler"
	"production_api/repository"
	"production_api/routes"
	"production_api/service"
)

func main() {
	// ---- Dependency Injection (Task 4) ----
	// Repository → Service → Handler (each depends on interface)
	repo := repository.NewInMemoryEmployeeRepo()   // Task 3: Repository layer
	svc := service.NewEmployeeService(repo)          // Task 2: Service layer
	empHandler := handler.NewEmployeeHandler(svc)    // Task 2: Handler layer

	// ---- Routes + Middleware ----
	router := routes.SetupRoutes(empHandler, svc)

	// ---- Start Server ----
	fmt.Println("╔════════════════════════════════════════════════════════════╗")
	fmt.Println("║   Day 12: Production REST API Architecture               ║")
	fmt.Println("║   Architecture: Handler → Service → Repository           ║")
	fmt.Println("║   No PostgreSQL — in-memory store with seed data         ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║   Public Endpoints:                                      ║")
	fmt.Println("║     GET  /health                     (Task 13)           ║")
	fmt.Println("║     GET  /ready                      (Task 14)           ║")
	fmt.Println("║   Protected Endpoints (Bearer token required):           ║")
	fmt.Println("║     GET  /api/employees              (Tasks 10,11,12)    ║")
	fmt.Println("║     GET  /api/employees/{id}                             ║")
	fmt.Println("║     POST /api/employees              (admin only)        ║")
	fmt.Println("║     PUT  /api/employees/{id}         (admin only)        ║")
	fmt.Println("║     DELETE /api/employees/{id}       (admin only)        ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║   Auth Tokens:                                           ║")
	fmt.Println("║     admin-token-123  → admin (read + write)              ║")
	fmt.Println("║     user-token-456   → user  (read only)                 ║")
	fmt.Println("╠════════════════════════════════════════════════════════════╣")
	fmt.Println("║   Query Params:                                          ║")
	fmt.Println("║     ?name=piyush&department=Engineering                   ║")
	fmt.Println("║     ?min_salary=50000&max_salary=80000                    ║")
	fmt.Println("║     ?sort_by=salary&sort_order=desc                       ║")
	fmt.Println("║     ?page=1&page_size=2                                   ║")
	fmt.Println("╚════════════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println("Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
╔══════════════════════════════════════════════════════╗
║                  CURL TEST COMMANDS                  ║
╚══════════════════════════════════════════════════════╝

# --- Public endpoints (no auth) ---

# Task 13: Health check
curl http://localhost:8080/health

# Task 14: Ready check
curl http://localhost:8080/ready


# --- Read endpoints (any token works) ---

# List all employees (Task 10: pagination, Task 11: filter, Task 12: sort)
curl -H "Authorization: Bearer user-token-456" http://localhost:8080/api/employees

# Pagination (Task 10)
curl -H "Authorization: Bearer user-token-456" "http://localhost:8080/api/employees?page=1&page_size=2"

# Filtering (Task 11)
curl -H "Authorization: Bearer user-token-456" "http://localhost:8080/api/employees?department=Engineering"
curl -H "Authorization: Bearer user-token-456" "http://localhost:8080/api/employees?name=piyush"
curl -H "Authorization: Bearer user-token-456" "http://localhost:8080/api/employees?min_salary=50000"

# Sorting (Task 12)
curl -H "Authorization: Bearer user-token-456" "http://localhost:8080/api/employees?sort_by=salary&sort_order=desc"

# Get by ID
curl -H "Authorization: Bearer user-token-456" http://localhost:8080/api/employees/1


# --- Write endpoints (admin token ONLY) ---

# Create (Task 7: auth, Task 8: admin, Task 9: validation)
curl -X POST http://localhost:8080/api/employees \
  -H "Authorization: Bearer admin-token-123" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Eve\",\"email\":\"eve@test.com\",\"age\":26,\"department\":\"Finance\",\"salary\":52000,\"role\":\"developer\"}"

# Update
curl -X PUT http://localhost:8080/api/employees/1 \
  -H "Authorization: Bearer admin-token-123" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Piyush Updated\",\"email\":\"piyush@new.com\",\"age\":26,\"department\":\"Backend\",\"salary\":75000,\"role\":\"lead\"}"

# Delete
curl -X DELETE http://localhost:8080/api/employees/3 \
  -H "Authorization: Bearer admin-token-123"


# --- Error cases ---

# No auth header (Task 7: 401)
curl http://localhost:8080/api/employees

# User tries to write (Task 8: 403)
curl -X POST http://localhost:8080/api/employees \
  -H "Authorization: Bearer user-token-456" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"Test\"}"

# Validation error (Task 9: 422)
curl -X POST http://localhost:8080/api/employees \
  -H "Authorization: Bearer admin-token-123" \
  -H "Content-Type: application/json" \
  -d "{\"name\":\"\",\"email\":\"bad\",\"age\":10}"

# Not found (404)
curl -H "Authorization: Bearer user-token-456" http://localhost:8080/api/employees/999
*/
