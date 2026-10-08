package main

import (
	"log"
	"net/http"

	"employee-management-go/controller"
	"employee-management-go/middleware"
	"employee-management-go/repository"
	"employee-management-go/service"

	"github.com/gorilla/mux"
)

func main() {

	// Create repository.
	employeeRepository := &repository.EmployeeRepositoryImpl{}

	// Create service.
	employeeService := &service.EmployeeServiceImpl{
		Repository: employeeRepository,
	}

	// Create controller.
	employeeController := controller.NewEmployeeController(employeeService)

	// Create router.
	router := mux.NewRouter()

	// Register REST APIs.
	router.HandleFunc(
		"/employees",
		employeeController.CreateEmployee,
	).Methods("POST")

	router.HandleFunc(
		"/employees",
		employeeController.GetEmployees,
	).Methods("GET")

	router.HandleFunc(
		"/employees/{id}",
		employeeController.GetEmployee,
	).Methods("GET")

	router.HandleFunc(
		"/employees/{id}",
		employeeController.DeleteEmployee,
	).Methods("DELETE")

	// Add logging middleware.
	router.Use(middleware.LoggingMiddleware)

	// Start server.
	log.Println("Server started on :8080")

	log.Fatal(
		http.ListenAndServe(":8080", router),
	)
}
