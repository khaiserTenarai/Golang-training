// Day 13 — Configuration & Application Lifecycle
// Task 15: Mini Project — Configurable Employee API with:
//   Environment configuration, Secrets, Logging, Feature flags, Graceful shutdown
//
// Tasks 9, 10: Graceful shutdown + SIGINT/SIGTERM handling
// Task 12: Startup initialization sequence
// Task 13: Shutdown sequence
//
// Run:  go mod tidy && go run main.go
// Override config: APP_ENV=production FEATURE_AUTH=true go run main.go
// Stop gracefully: press Ctrl+C (SIGINT)

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"configurable_api/config"
	"configurable_api/features"
	"configurable_api/handler"
	"configurable_api/middleware"
	"configurable_api/repository"
	"configurable_api/service"

	"github.com/gorilla/mux"
)

func main() {
	// ================================================================
	// Task 12: Startup Initialization Sequence
	// ================================================================

	fmt.Println("╔══════════════════════════════════════════════════╗")
	fmt.Println("║  Day 13: Configuration & Application Lifecycle  ║")
	fmt.Println("╚══════════════════════════════════════════════════╝")
	fmt.Println()

	// Step 1: Load configuration (Tasks 1-4)
	fmt.Println("[STARTUP] Step 1: Loading configuration...")
	cfg := config.Load()

	// Step 2: Validate configuration (Task 3)
	fmt.Println("[STARTUP] Step 2: Validating configuration...")
	if err := cfg.Validate(); err != nil {
		log.Fatalf("[FATAL] %v", err)
	}
	cfg.Print()

	// Step 3: Load feature flags (Task 14)
	fmt.Println("[STARTUP] Step 3: Loading feature flags...")
	flags := features.LoadFlags(cfg)
	fmt.Printf("  Features: pagination=%v filtering=%v sorting=%v auth=%v\n",
		flags.Pagination, flags.Filtering, flags.Sorting, flags.Auth)

	// Step 4: Initialize repository
	fmt.Println("[STARTUP] Step 4: Initializing repository (in-memory)...")
	repo := repository.NewInMemoryRepo()

	// Step 5: Initialize service (Task 11: context-aware)
	fmt.Println("[STARTUP] Step 5: Initializing service layer...")
	svc := service.NewEmployeeService(repo)

	// Step 6: Initialize handlers
	fmt.Println("[STARTUP] Step 6: Initializing handlers...")
	empHandler := handler.NewEmployeeHandler(svc, flags, cfg.MaxPageSize)

	// Step 7: Setup routes + middleware
	fmt.Println("[STARTUP] Step 7: Setting up routes and middleware...")
	router := mux.NewRouter()
	router.Use(middleware.Recovery)
	router.Use(middleware.RequestID) // Task 5
	router.Use(middleware.Logging)   // Task 6

	// Public endpoints (Tasks 13, 14)
	router.HandleFunc("/health", empHandler.Health).Methods("GET")
	router.HandleFunc("/ready", empHandler.Ready).Methods("GET")

	// API routes
	api := router.PathPrefix("/api").Subrouter()
	if flags.Auth { // Task 14: feature flag controls auth
		api.Use(middleware.Auth(cfg.Secrets.AdminToken)) // Task 7/8
		fmt.Println("  [AUTH] Authentication enabled")
	} else {
		fmt.Println("  [AUTH] Authentication disabled (dev mode)")
	}

	api.HandleFunc("/employees", empHandler.GetAll).Methods("GET")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.GetByID).Methods("GET")
	api.HandleFunc("/employees", empHandler.Create).Methods("POST")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.Update).Methods("PUT")
	api.HandleFunc("/employees/{id:[0-9]+}", empHandler.Delete).Methods("DELETE")

	// Step 8: Create HTTP server
	server := &http.Server{
		Addr:         cfg.ServerHost + ":" + cfg.ServerPort,
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSecs) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSecs) * time.Second,
	}

	// ================================================================
	// Tasks 9, 10: Graceful Shutdown + Signal Handling
	// ================================================================

	// Task 10: Listen for SIGINT (Ctrl+C) and SIGTERM
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start server in background
	go func() {
		fmt.Printf("\n[STARTUP] ✓ Server running on http://%s:%s (%s)\n\n",
			cfg.ServerHost, cfg.ServerPort, cfg.Environment)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("[FATAL] Server error: %v", err)
		}
	}()

	// Block until signal received
	sig := <-sigChan
	fmt.Printf("\n[SHUTDOWN] Signal received: %v\n", sig)

	// ================================================================
	// Task 13: Shutdown Sequence
	// ================================================================

	// Task 9: Graceful shutdown with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		time.Duration(cfg.ShutdownTimeoutSecs)*time.Second,
	)
	defer shutdownCancel()

	fmt.Printf("[SHUTDOWN] Step 1: Draining connections (timeout: %ds)...\n", cfg.ShutdownTimeoutSecs)
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("[SHUTDOWN] Forced shutdown: %v", err)
	} else {
		fmt.Println("[SHUTDOWN] Step 2: HTTP server stopped gracefully")
	}

	fmt.Println("[SHUTDOWN] Step 3: Flushing logs...")
	time.Sleep(100 * time.Millisecond) // simulate log flush

	fmt.Println("[SHUTDOWN] Step 4: Cleanup complete")
	fmt.Println("[SHUTDOWN] ✓ Application exited cleanly")
}

/*
══════════════════════════════════════
  CURL TEST COMMANDS
══════════════════════════════════════

# Health/Ready (always public)
curl http://localhost:8080/health
curl http://localhost:8080/ready

# CRUD (no auth in dev mode)
curl http://localhost:8080/api/employees
curl http://localhost:8080/api/employees/1
curl -X POST http://localhost:8080/api/employees -H "Content-Type: application/json" -d "{\"name\":\"Eve\",\"email\":\"eve@test.com\",\"age\":26,\"department\":\"HR\",\"salary\":52000}"
curl -X PUT http://localhost:8080/api/employees/1 -H "Content-Type: application/json" -d "{\"name\":\"Piyush Updated\",\"email\":\"piyush@new.com\",\"age\":26,\"department\":\"Backend\",\"salary\":70000}"
curl -X DELETE http://localhost:8080/api/employees/3

# Filtering + Sorting + Pagination
curl "http://localhost:8080/api/employees?department=Engineering&sort_by=salary&sort_order=desc&page=1&page_size=2"

# Run with production config
APP_ENV=production FEATURE_AUTH=true go run main.go
# Then use token:
curl -H "Authorization: Bearer admin-secret-token-123" http://localhost:8080/api/employees

# Graceful shutdown: press Ctrl+C and watch the shutdown sequence
*/
