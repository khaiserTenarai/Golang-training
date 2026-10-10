// Package main starts the Gin Employee Management API.
package main

import (
	"context"
	"employee-management/config"
	"employee-management/controller"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Load database and server configuration from environment variables.
	cfg := config.Load()

	// Create a logger to record application messages and errors.
	logger := logging.New()

	// Build the database connection string using configuration values.
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Create a context with a 10-second timeout for database operations.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Release the context resources when main finishes.
	defer cancel()

	// Create a connection pool for PostgreSQL.
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		// Log the error and stop if the connection pool cannot be created.
		logger.Error(err.Error())
		return
	}

	// Close all database connections when main finishes.
	defer db.Close()

	// Check whether PostgreSQL is reachable.
	if err := db.Ping(ctx); err != nil {
		// Log the error and stop if the database is unavailable.
		logger.Error(err.Error())
		return
	}

	// Create the repository to perform employee database operations.
	repo := repository.NewPostgresEmployeeRepository(db)

	// Create the service to handle employee business logic.
	svc := service.NewEmployeeService(repo)

	// Create the controller to handle incoming employee requests.
	ctl := controller.NewEmployeeController(svc)

	// Create a new Gin router without default middleware.
	r := gin.New()

	// Recover from request-handling panics and log HTTP requests.
	r.Use(gin.Recovery(), middleware.Logging(logger))

	// Register employee API routes with the Gin router.
	ctl.RegisterRoutes(r)

	// Configure the HTTP server and its timeout settings.
	server := &http.Server{
		Addr:         ":" + cfg.ServerPort, // Port on which the server listens.
		Handler:      r,                    // Gin router handles incoming requests.
		ReadTimeout:  10 * time.Second,     // Maximum time to read a request.
		WriteTimeout: 10 * time.Second,     // Maximum time to write a response.
		IdleTimeout:  60 * time.Second,     // Maximum time to keep an idle connection open.
	}

	// Log the address where the API is running.
	logger.Info("Gin Employee Management API started on http://localhost:" + cfg.ServerPort)

	// Start the HTTP server and listen for incoming requests.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		// Log the error if the server stops unexpectedly.
		logger.Error(err.Error())
	}
}

/*






























// -----------------------------------

// Package main starts the Gin Employee Management API.

// @title Employee Management API
// @version 1.0
// @description REST API for managing employees.
// @host localhost:8080
// @BasePath /

package main

import (
	"context"
	"employee-management/config"
	"employee-management/controller"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	// Import generated Swagger documentation.
	_ "employee-management/docs"

	// Swagger UI packages.
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load application configuration.
	cfg := config.Load()

	// Create the application logger.
	logger := logging.New()

	// Build the PostgreSQL connection string.
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	// Create a context with a 10-second timeout for database setup.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create the PostgreSQL connection pool.
	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.Close()

	// Verify that PostgreSQL is reachable.
	if err := db.Ping(ctx); err != nil {
		logger.Error(err.Error())
		return
	}

	// Initialize repository, service, and controller layers.
	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctl := controller.NewEmployeeController(svc)

	// Create the Gin router.
	r := gin.New()

	// Register recovery and logging middleware.
	r.Use(gin.Recovery(), middleware.Logging(logger))

	// Register employee CRUD routes.
	ctl.RegisterRoutes(r)

	// Register Swagger UI.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configure the HTTP server.
	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the API server.
	logger.Info("Gin Employee Management API started on http://localhost:" + cfg.ServerPort)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}

*/
