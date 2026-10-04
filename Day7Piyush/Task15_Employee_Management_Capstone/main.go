package main

import (
	"fmt"
	"log"
	"net/http"
	"task15_employee_management_capstone/config"
	"task15_employee_management_capstone/controller"
	"task15_employee_management_capstone/repository"
	"task15_employee_management_capstone/routes"
	"task15_employee_management_capstone/service"
)

func main() {
	// ---- Database ----
	db := config.ConnectDB()
	defer db.Close()
	config.CreateTables(db)

	// ---- Repository Layer (interfaces) ----
	userRepo := repository.NewPostgresUserRepository(db)
	empRepo := repository.NewPostgresEmployeeRepository(db)
	deptRepo := repository.NewPostgresDepartmentRepository(db)

	// ---- Service Layer (business logic, depends on repo interfaces) ----
	authSvc := service.NewAuthService(userRepo)
	empSvc := service.NewEmployeeService(empRepo)
	deptSvc := service.NewDepartmentService(deptRepo)

	// ---- Controller Layer (HTTP handlers, depends on service interfaces) ----
	authCtrl := controller.NewAuthController(authSvc)
	empCtrl := controller.NewEmployeeController(empSvc)
	deptCtrl := controller.NewDepartmentController(deptSvc)

	// ---- Routes + Middleware ----
	router := routes.SetupRoutes(authCtrl, empCtrl, deptCtrl, authSvc)

	fmt.Println("=== EMPLOYEE MANAGEMENT CAPSTONE ===")
	fmt.Println("Architecture: Controller -> Service -> Repository -> PostgreSQL")
	fmt.Println("Features: REST API, CRUD, Auth, Pagination, Transactions, Logging")
	fmt.Println()
	fmt.Println("Server running on http://localhost:8084")
	fmt.Println()
	fmt.Println("Public:")
	fmt.Println("  POST /api/auth/register")
	fmt.Println("  POST /api/auth/login")
	fmt.Println("Protected (Bearer token):")
	fmt.Println("  POST/GET        /api/employees")
	fmt.Println("  GET/PUT/DELETE   /api/employees/{id}")
	fmt.Println("  PUT              /api/employees/{id}/salary")
	fmt.Println("  GET              /api/employees/{id}/salary-history")
	fmt.Println("  POST/GET        /api/departments")
	fmt.Println("  GET/PUT/DELETE   /api/departments/{id}")
	fmt.Println("  GET  /api/auth/profile")
	fmt.Println("  POST /api/auth/logout")
	log.Fatal(http.ListenAndServe(":8084", router))
}
