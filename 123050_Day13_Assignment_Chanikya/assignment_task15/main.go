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

	"day13_project/config"
	"day13_project/controller"
	"day13_project/middleware"
	"day13_project/repository"
	"day13_project/service"
)

func main() {
	log.Println("[STARTUP] Initializing sequence config loaders...")
	
	cfg, _, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Fatal system instantiation configuration parameter layout failure: %v", err)
	}

	// Startup Initialization Sequence
	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo, cfg)
	ctrl := controller.NewEmployeeController(svc)

	mux := http.NewServeMux()
	mux.HandleFunc("/employees", ctrl.GetEmployees)
	mux.HandleFunc("/employees/create", ctrl.CreateEmployee)
	mux.HandleFunc("/health", ctrl.HealthCheck)
	mux.HandleFunc("/ready", ctrl.ReadinessCheck)

	handlerChain := middleware.TelemetryAndSecurityChain(mux)

	server := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: handlerChain,
	}

	go func() {
		log.Printf("[STARTUP] Server active. Listening on infrastructure port: %s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server binding allocation fault error: %v", err)
		}
	}()

	// Graceful Shutdown & Signal Handlers
	shutdownSig := make(chan os.Signal, 1)
	signal.Notify(shutdownSig, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	<-shutdownSig
	log.Println("[SHUTDOWN] Intercept system signal initialized. Gracefully running shutdown sequence...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("[SHUTDOWN] Force closing context instance loop. Fault trace: %v", err)
	}

	log.Println("[SHUTDOWN] Server termination sequence executed safely. System offline.")
}
