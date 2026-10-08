package main

import (
	"fmt"
	"net/http"
	"production-employee-api/handlers"
	"production-employee-api/middleware"
	"production-employee-api/repository"
	"production-employee-api/service"

	
)

func main() {
	// Dependency Injection Setup
	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo)
	handler := handlers.NewEmployeeHandler(svc)

	mux := http.NewServeMux()

	// Public Health Endpoints
	mux.HandleFunc("/health", handler.HealthCheck)
	mux.HandleFunc("/ready", handler.ReadyCheck)

	// Protected Employee Routes
	mux.Handle("/api/v1/employees", middleware.Authenticate(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetEmployees(w, r)
		} else if r.Method == http.MethodPost {
			middleware.AuthorizeRole("admin", handler.CreateEmployee)(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})))

	// Apply Middlewares
	handlerStack := middleware.RequestID(middleware.Logger(mux))

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", handlerStack); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}