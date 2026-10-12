package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"employee-api/config"
	"employee-api/internal/controller"
	"employee-api/internal/database"
	"employee-api/internal/repository"
	"employee-api/internal/server"
	"employee-api/internal/service"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("[STARTUP] Step 1: Loading Configuration...")
	cfg, err := config.LoadConfig("config")
	if err != nil {
		log.Fatalf("[FATAL] Application startup failed: %v", err)
	}

	log.Println("[STARTUP] Step 2: Connecting to Database...")
	db, err := database.NewPostgresDB(ctx, cfg.Database)
	if err != nil {
		log.Fatalf("[FATAL] Database connection failed: %v", err)
	}
	defer db.Close()

	log.Println("[STARTUP] Step 3: Wiring Application Layers...")
	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)

	router := server.NewRouter(ctrl)
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("[SERVER] Listening on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErrors <- err
		}
	}()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("[FATAL] HTTP Server error: %v", err)
	case sig := <-shutdownSignal:
		log.Printf("[LIFECYCLE] Shutdown signal received: %v. Initiating graceful shutdown...", sig)
		
		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
		defer shutdownCancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Printf("[SHUTDOWN] HTTP server force shutdown error: %v", err)
			_ = srv.Close()
		}
		log.Println("[SHUTDOWN] Graceful shutdown completed cleanly.")
	}
}
