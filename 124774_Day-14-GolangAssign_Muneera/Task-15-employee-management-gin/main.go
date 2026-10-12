// // -----------------------------------

// Package main starts the Gin Employee Management API.

// @title Employee Management API
// @version 1.0
// @description REST API for managing employees.
// @host localhost:10007
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
