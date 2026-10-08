// Package main is the entry point of the Employee Management REST API.
package main

// Import standard packages used by the application.
import (
	// Import context for database ping and lifecycle operations.
	"context"
	// Import fmt for building the PostgreSQL connection string.
	"fmt"
	// Import net/http for the REST API server.
	"net/http"
	// Import time for database connection timeout configuration.
	"time"
)

// Import pgxpool for PostgreSQL connection pooling.
import "github.com/jackc/pgx/v5/pgxpool"

// Import application configuration.
import "employee-management/config"

// Import employee controller.
import "employee-management/controller"

// Import logging package.
import "employee-management/logging"

// Import logging middleware.
import "employee-management/middleware"

// Import employee repository.
import "employee-management/repository"

// Import employee service.
import "employee-management/service"

// main starts the application.
func main() {
	// Load configuration from .env and environment variables.
	cfg := config.Load()
	// Create the application logger.
	logger := logging.New()
	// Build the PostgreSQL connection string.
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	// Create a context with a timeout for the database connection.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Ensure the context resources are released.
	defer cancel()
	// Create a PostgreSQL connection pool.
	db, err := pgxpool.New(ctx, dsn)
	// Stop the application when the connection pool cannot be created.
	if err != nil {
		// Log the database initialization error.
		logger.Error(err.Error())
		// Exit the application.
		return
	}
	// Ensure the database pool is closed during application shutdown.
	defer db.Close()
	// Verify that PostgreSQL is reachable.
	if err := db.Ping(ctx); err != nil {
		// Log the PostgreSQL connection error.
		logger.Error(err.Error())
		// Exit the application.
		return
	}
	// Log successful PostgreSQL connectivity.
	logger.Info("connected to PostgreSQL")
	// Create the repository implementation.
	employeeRepository := repository.NewPostgresEmployeeRepository(db)
	// Create the service using repository dependency injection.
	employeeService := service.NewEmployeeService(employeeRepository)
	// Create the HTTP controller using service dependency injection.
	employeeController := controller.NewEmployeeController(employeeService)
	// Create the standard Go HTTP router.
	mux := http.NewServeMux()
	// Register employee REST endpoints.
	employeeController.RegisterRoutes(mux)
	// Create the logging middleware around the router.
	handler := middleware.Logging(mux, logger)
	// Build the HTTP server.
	server := &http.Server{
		// Configure the server address from the .env file.
		Addr: ":" + cfg.ServerPort,
		// Configure the application handler.
		Handler: handler,
		// Configure an HTTP read timeout.
		ReadTimeout: 10 * time.Second,
		// Configure an HTTP write timeout.
		WriteTimeout: 10 * time.Second,
		// Configure an HTTP idle timeout.
		IdleTimeout: 60 * time.Second,
	}
	// Log the server startup message.
	logger.Info("Employee Management API started on http://localhost:" + cfg.ServerPort)
	// Start the HTTP server.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Log unexpected server errors.
		logger.Error(err.Error())
	}
}
