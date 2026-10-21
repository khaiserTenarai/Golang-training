package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"employee_api/config"
	"employee_api/controller"
	"employee_api/middleware"
	"employee_api/repository"
	"employee_api/service"
)

func main() {
	log.Println("Starting Employee API Server initialization...")

	// 1. Load Configurations
	cfg := config.LoadConfig()

	// 2. Initialize Layers (Dependency Injection)
	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo, cfg)
	ctrl := controller.NewEmployeeController(svc)

	// 3. Setup Routes and Middleware Chains
	mux := http.NewServeMux()

	// Health and Readiness probes
	mux.HandleFunc("/health", ctrl.LivenessProbe)
	mux.HandleFunc("/ready", ctrl.ReadinessProbe)

	// Protected business logic API endpoints
	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ctrl.ListEmployees(w, r)
		case http.MethodPost:
			ctrl.CreateEmployee(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Inject Middleware Chain: Request ID -> Telemetry Logging -> AuthN -> AuthZ (Admin only for safety demo)
	protectedChain := middleware.RequestID(
		middleware.Logger(
			middleware.AuthN(
				middleware.AuthZ("admin", apiMux),
			),
		),
	)

	mux.Handle("/employees", protectedChain)

	// 4. Start HTTP Server
	serverAddr := fmt.Sprintf(":%s", cfg.ServerPort)
	srv := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	// Channel to catch termination signals for graceful shutdown
	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("Server responding on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server startup failed: %v", err)
		}
	}()

	// 5. Graceful Shutdown Block
	sig := <-shutdownChan
	log.Printf("Received system signal: %v. Initiating graceful shutdown sequence...", sig)

	// Allow maximum 5 seconds cleanup context window
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown violently: %v", err)
	}

	log.Println("Employee API Server finished clean termination sequence. Goodbye.")
}
