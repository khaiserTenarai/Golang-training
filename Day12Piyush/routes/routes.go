// Route setup — wires handlers with middleware.
// Task 1: Layered architecture visible here in the middleware chain.
// Tasks 5-8: Middleware applied per route group.

package routes

import (
	"production_api/handler"
	"production_api/middleware"
	"production_api/service"

	"github.com/gorilla/mux"
)

func SetupRoutes(empHandler *handler.EmployeeHandler, empService service.EmployeeService) *mux.Router {
	r := mux.NewRouter()

	// Global middleware (applied to ALL routes)
	r.Use(middleware.RecoveryMiddleware)   // Panic recovery
	r.Use(middleware.RequestIDMiddleware)  // Task 5: Request-ID
	r.Use(middleware.LoggingMiddleware)    // Task 6: Logging

	// Task 13: /health — public, no auth needed
	r.HandleFunc("/health", handler.HealthHandler).Methods("GET")

	// Task 14: /ready — public, no auth needed
	r.HandleFunc("/ready", handler.ReadyHandler(empService)).Methods("GET")

	// Protected API routes (Task 7: auth required)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)    // Task 7: Authentication
	api.Use(middleware.AdminWriteOnly)    // Task 8: Authorization

	// Employee CRUD
	api.HandleFunc("/employees", empHandler.GetAll).Methods("GET")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.GetByID).Methods("GET")
	api.HandleFunc("/employees", empHandler.Create).Methods("POST")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.Update).Methods("PUT")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.Delete).Methods("DELETE")

	return r
}
