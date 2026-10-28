package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"employee-api/config"
	"employee-api/database"
	"employee-api/server"
)

func main() {
	// Run the application.
	// Exit with an error code if startup or shutdown fails.
	if err := run(); err != nil {
		log.Printf("Application stopped with error: %v", err)
		os.Exit(1)
	}
}

// run manages application startup, execution, and cleanup.
func run() error {
	// Step 1: Load normal application configuration.
	cfg := config.Load()

	// Step 2: Check that configuration values are valid.
	if err := cfg.Validate(); err != nil {
		return err
	}

	// Step 3: Load the database password separately.
	secrets, err := config.LoadSecrets()
	if err != nil {
		return err
	}

	// Step 4: Connect to PostgreSQL using GORM.
	db, err := database.Connect(cfg, secrets)
	if err != nil {
		return err
	}

	// Step 5: Get the underlying database connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Step 6: Close the database pool when run() returns.
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed to close database pool: %v", err)
		} else {
			log.Println("Database connection pool closed")
		}
	}()

	// Step 7: Create the HTTP server and its routes.
	httpServer := server.New(cfg, db)

	// Step 8: Listen for Ctrl+C or a SIGTERM signal.
	// When a signal arrives, the context is cancelled.
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Step 9: Print startup information.
	log.Printf(
		"Starting %s server on port %s",
		cfg.AppEnv,
		cfg.AppPort,
	)
	log.Println("Health endpoint: /health")
	log.Println("Database health endpoint: /health/db")

	// Step 10: Run the server until it stops or shutdown is requested.
	if err := server.Start(ctx, httpServer); err != nil {
		return err
	}

	// Step 11: Report that the HTTP server has stopped.
	log.Println("HTTP server shut down successfully")

	// Returning from run() triggers the deferred database cleanup.
	return nil
}
