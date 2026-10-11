// Package main is the entry point of the Employee Management REST API.
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/config"
	"employee-management/controller"
	"employee-management/logging"
	"employee-management/middleware"
	"employee-management/repository"
	"employee-management/service"
)

func main() {
	cfg := config.Load()
	logger := logging.New()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, dsn)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("connected to PostgreSQL")

	// Table auto-creation script
	_, _ = db.Exec(ctx, `CREATE TABLE IF NOT EXISTS employees (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		department VARCHAR(50) NOT NULL,
		salary NUMERIC(10,2) NOT NULL
	);`)

	employeeRepository := repository.NewPostgresEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)

	mux := http.NewServeMux()
	employeeController.RegisterRoutes(mux)

	// Combine RequestID and Logging middleware layers
	handler := middleware.RequestID(middleware.Logging(mux, logger))

	server := &http.Server{
		Addr:         ":" + cfg.ServerPort,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	logger.Info("Employee Management API started on http://localhost:" + cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}