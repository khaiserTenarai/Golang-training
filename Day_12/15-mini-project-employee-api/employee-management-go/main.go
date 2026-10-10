// Package main starts the Employee Management REST API.
package main

import (
	"context"
	"employee-management/auth"
	"employee-management/config"
	"employee-management/controller"
	"employee-management/health"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// main initializes dependencies, routes and the HTTP server.

func main() {
	// Call Load() function that  loads configuration from .env and environment variables.
	// store configuration in struct
	// cfg contains database details, server port, JWT details and user credentials.
	cfg := config.Load()

	// Create the application logger.
	// The logger is used to print information and errors.
	logger := logging.New()

	// Create the PostgreSQL connection string using the database details from cfg.
	// The DSN tells Go how to connect to PostgreSQL.
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	// Create a context with a timeout of 10 seconds.
	// This prevents the database connection from waiting forever.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Release the context resources when main finishes.
	defer cancel()

	// Create a connection pool for PostgreSQL.
	// The pool manages multiple database connections.
	db, err := pgxpool.New(ctx, dsn)

	// Check whether creating the database connection pool failed.
	if err != nil {
		// Log the database error.
		logger.Error(err.Error())

		// Stop the application because the database is required.
		return
	}

	// Close the database connection pool when the application stops.
	defer db.Close()

	// Check whether the application can actually communicate with PostgreSQL.
	if err := db.Ping(ctx); err != nil {
		// Log the database connection error.
		logger.Error(err.Error())

		// Stop the application if the database is not available.
		return
	}

	// Create the repository.
	// The repository is responsible for database operations.
	repo := repository.NewPostgresEmployeeRepository(db)

	// Create the service.
	// The service contains employee business logic.
	// It receives the repository as a dependency.
	svc := service.NewEmployeeService(repo)

	// Create the employee controller.
	// The controller handles HTTP requests related to employees.
	// It receives the service as a dependency.
	employeeController := controller.NewEmployeeController(svc)
	// ---------------------------

	// Create the authentication service.
	// It contains JWT and login-related logic.
	// JWT secret, expiry time and user credentials come from configuration.
	authService := auth.New(cfg.JWTSecret, cfg.JWTExpiryMinutes, cfg.AdminUsername, cfg.AdminPasswordHash, cfg.UserUsername, cfg.UserPasswordHash)

	// Create the authentication controller.
	// This controller handles login requests.
	authController := controller.NewAuthController(authService)

	healthHandler := health.NewHandler(svc)

	// Create the main HTTP router
	// it recieves inoming HTTP requests
	// The router matches the incoming URL with the correct handler.meaning which handler should handle each request..
	mux := http.NewServeMux()

	// Register the login API.
	// Login is public because the user does not have a JWT yet.
	mux.HandleFunc("/auth/login", authController.Login)

	mux.HandleFunc("/health", healthHandler.Health)
	mux.HandleFunc("/ready", healthHandler.Ready)
	mux.HandleFunc("/docs/openapi.json", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "docs/openapi.json") })
	mux.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "web/swagger.html") })

	// Public login, health and documentation routes remain open; employee routes are registered through protected middleware.

	// Create a separate router for employee APIs.
	employeeMux := http.NewServeMux()

	// Register all employee-related routes in the employee router.
	employeeController.RegisterRoutes(employeeMux)

	// Add authentication middleware to employee routes.
	// The JWT is checked before the request reaches the employee handler.
	// AdminWrites adds authorization rules for write operations.
	protectedEmployees := middleware.RequireAuth(authService, middleware.AdminWrites(employeeMux))

	// Connect the /employees route to the protected employee handler.
	// Every request to /employees must pass through authentication.
	mux.Handle("/employees", protectedEmployees)

	// Connect routes such as /employees/1 to the protected employee handler.
	// The request must also pass authentication and authorization checks.
	mux.Handle("/employees/", protectedEmployees)

	// Add logging middleware around the complete main router.
	// This logs information about incoming HTTP requests.
	handler := middleware.Logging(mux, logger)

	// Create the HTTP server.
	// Addr specifies the port.
	// Handler specifies the main request handler.
	// ReadTimeout limits how long the server waits to read a request.
	// WriteTimeout limits how long the server waits to write a response.
	// IdleTimeout controls how long an idle connection can remain open.
	server := &http.Server{Addr: ":" + cfg.ServerPort, Handler: handler, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}

	// Log a message showing that the API server has started.
	logger.Info("Employee Management API started on http://localhost:" + cfg.ServerPort)

	// Start listening for HTTP requests.
	// Postman/browser/client requests will come to this server.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Log the server error if the server stops unexpectedly.
		logger.Error(err.Error())
	}
}
