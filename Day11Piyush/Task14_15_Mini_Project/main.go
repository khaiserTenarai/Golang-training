// Tasks 14, 15: Day 11 Mini Project — PostgreSQL-backed Employee REST API
// Task 14 — Connect the application to PostgreSQL
// Task 15 — Complete CRUD REST API with PostgreSQL
//
// Prerequisites:
//   CREATE DATABASE day11db;
//
// Run:
//   go mod tidy
//   go run main.go
//
// Edit config/database.go if your PostgreSQL credentials differ.

package main

import (
	"fmt"
	"log"
	"net/http"
	"task14_15_employee_api/config"
	"task14_15_employee_api/controller"
	"task14_15_employee_api/repository"
	"task14_15_employee_api/routes"
)

func main() {
	// Task 14: Connect to PostgreSQL
	db := config.ConnectDB()
	defer db.Close()
	config.CreateTable(db)

	// Task 15: Wire up layers
	repo := repository.NewEmployeeRepository(db)
	ctrl := controller.NewEmployeeController(repo)
	router := routes.SetupRoutes(ctrl)

	fmt.Println("\n===== Day 11 Mini Project: PostgreSQL Employee REST API =====")
	fmt.Println("Server running on http://localhost:8081")
	fmt.Println()
	fmt.Println("Endpoints:")
	fmt.Println("  GET    /api/employees                             → List (+ query params)")
	fmt.Println("  GET    /api/employees/{id}                        → Detail")
	fmt.Println("  POST   /api/employees                             → Create")
	fmt.Println("  PUT    /api/employees/{id}                        → Update")
	fmt.Println("  DELETE /api/employees/{id}                        → Delete")
	fmt.Println()
	fmt.Println("Query params: ?name=, ?department=, ?min_salary=, ?max_salary=, ?page=, ?page_size=")
	fmt.Println()
	fmt.Println("Test:")
	fmt.Println(`  curl -X POST http://localhost:8081/api/employees -H "Content-Type: application/json" -d "{\"name\":\"Piyush\",\"email\":\"piyush@test.com\",\"age\":25,\"department\":\"Engineering\",\"salary\":60000}"`)
	fmt.Println(`  curl http://localhost:8081/api/employees`)
	fmt.Println(`  curl http://localhost:8081/api/employees/1`)
	log.Fatal(http.ListenAndServe(":8081", router))
}
