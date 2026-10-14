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
	cfg := config.Load()
	logger := logging.New()
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
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
	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	employeeController := controller.NewEmployeeController(svc)
	authService := auth.New(cfg.JWTSecret, cfg.JWTExpiryMinutes, cfg.AdminUsername, cfg.AdminPasswordHash, cfg.UserUsername, cfg.UserPasswordHash)
	authController := controller.NewAuthController(authService)
	healthHandler := health.NewHandler(svc)
	mux := http.NewServeMux()
	mux.HandleFunc("/auth/login", authController.Login)
	mux.HandleFunc("/health", healthHandler.Health)
	mux.HandleFunc("/ready", healthHandler.Ready)
	employeeMux := http.NewServeMux()
	employeeController.RegisterRoutes(employeeMux)
	protectedEmployees := middleware.RequireAuth(authService, middleware.AdminWrites(employeeMux))
	mux.Handle("/employees", protectedEmployees)
	mux.Handle("/employees/", protectedEmployees)
	handler := middleware.Logging(mux, logger)
	server := &http.Server{Addr: ":" + cfg.ServerPort, Handler: handler, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second, IdleTimeout: 60 * time.Second}
	logger.Info("Employee Management API started on http://localhost:" + cfg.ServerPort)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
	}
}
