package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"employee-api-day13/config"
	"employee-api-day13/controller"
	"employee-api-day13/database"
	"employee-api-day13/logger"
	"employee-api-day13/middleware"
	"employee-api-day13/repository"
	"employee-api-day13/service"
)

func main() {
	// Task 12: STARTUP SEQUENCE — load config, validate, initialize logger,
	// connect DB, initialize repository/service/controller, then start server.
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	appLogger := logger.New(cfg.LogLevel, cfg.AppEnv == "production")
	appLogger.Info("starting Employee API", "environment", cfg.AppEnv)

	startupCtx, cancelStartup := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	db, err := database.Open(startupCtx, cfg)
	cancelStartup()

	if err != nil {
		appLogger.Error("database initialization failed", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	repo := repository.NewPostgresEmployeeRepository(db)
	employeeService := service.NewEmployeeService(repo, cfg.FeatureFlags.CreateEmployee)
	employeeController := controller.NewEmployeeController(employeeService)

	mux := http.NewServeMux()
	employeeController.Register(mux)
	handler := middleware.Logging(appLogger, mux)

	server := &http.Server{
		Addr:              ":" + cfg.AppPort,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Task 10: Handle (SIGINT) and termination (SIGTERM).
	signalCtx, stopSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stopSignals()

	serverErrors := make(chan error, 1)
	go func() {
		appLogger.Info("HTTP server listening", "port", cfg.AppPort)
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if !errors.Is(err, http.ErrServerClosed) {
			appLogger.Error("HTTP server failed", "error", err)
		}
	case <-signalCtx.Done():
		appLogger.Info("shutdown signal received")
	}

	// Task 9 and Task 13: graceful shutdown sequence. Stop accepting new
	// requests, wait for in-flight requests, then close the PostgreSQL pool.
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()
	if err := server.Shutdown(shutdownCtx); err != nil {
		appLogger.Error("graceful shutdown failed; forcing close", "error", err)
		if closeErr := server.Close(); closeErr != nil {
			appLogger.Error("forced server close failed", "error", closeErr)
		}
	}
	db.Close()
	appLogger.Info("shutdown complete")

}
