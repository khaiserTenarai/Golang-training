package main

import (
	"context"
	"configurable-employee-api/config"
	"configurable-employee-api/handlers"
	"configurable-employee-api/repository"
	"configurable-employee-api/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 1. Startup: Load and validate configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[STARTUP ERROR] Configuration failed: %v", err)
	}

	log.Printf("[STARTUP] Environment: %s | Port: %d | Feature Flag: %v",
		cfg.AppEnv, cfg.Port, cfg.EnableNewFeature)

	// 2. Startup: Initialize layers (Dependency Injection)
	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo)
	handler := handlers.NewHandler(svc, cfg)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/employees", handler.GetEmployees)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	// 3. Graceful Shutdown listener setup
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("[STARTUP] Server starting on http://localhost:%d", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[SERVER ERROR] %v", err)
		}
	}()

	// Block until SIGINT or SIGTERM is received
	sig := <-shutdown
	log.Printf("[SHUTDOWN] Signal received (%v). Starting graceful shutdown...", sig)

	// 4. Shutdown Sequence
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("[SHUTDOWN ERROR] Server forced to shutdown: %v", err)
	} else {
		log.Println("[SHUTDOWN] HTTP server stopped cleanly.")
	}

	if err := repo.Close(); err != nil {
		log.Printf("[SHUTDOWN ERROR] Repository close error: %v", err)
	}

	log.Println("[SHUTDOWN] Application lifecycle completed.")
}