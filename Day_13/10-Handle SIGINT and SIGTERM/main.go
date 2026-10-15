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
	// Run the application and report startup or shutdown errors.
	if err := run(); err != nil {
		log.Printf("Application stopped with error: %v", err)
		os.Exit(1)
	}
}

// run manages startup, signal handling, and cleanup.
func run() error {
	// Step 1: Load configuration.
	cfg := config.Load()

	// Step 2: Validate configuration.
	if err := cfg.Validate(); err != nil {
		return err
	}

	// Step 3: Load the database password.
	secrets, err := config.LoadSecrets()
	if err != nil {
		return err
	}

	// Step 4: Connect to PostgreSQL.
	db, err := database.Connect(cfg, secrets)
	if err != nil {
		return err
	}

	// Step 5: Get the underlying SQL connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// Close the database pool when run() returns.
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed to close database pool: %v", err)
		} else {
			log.Println("Database connection pool closed")
		}
	}()

	// Step 6: Create the HTTP server.
	httpServer := server.New(cfg, db)

	// Step 7: Create a channel for operating-system signals.
	signalChannel := make(chan os.Signal, 1)

	// Receive Ctrl+C (SIGINT) and SIGTERM notifications.
	signal.Notify(
		signalChannel,
		os.Interrupt,
		syscall.SIGTERM,
	)

	// Stop receiving signal notifications when run() returns.
	defer signal.Stop(signalChannel)

	// Step 8: Create a context that controls server shutdown.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Step 9: Wait for a shutdown signal in a goroutine.
	go func() {
		receivedSignal := <-signalChannel

		log.Printf("Received signal: %s", receivedSignal)

		// Tell the server to begin graceful shutdown.
		cancel()
	}()

	// Step 10: Log startup information.
	log.Printf(
		"Starting %s server on port %s",
		cfg.AppEnv,
		cfg.AppPort,
	)
	log.Println("Health endpoint: /health")
	log.Println("Database health endpoint: /health/db")

	// Step 11: Run until the server stops or a signal arrives.
	if err := server.Start(ctx, httpServer); err != nil {
		return err
	}

	log.Println("HTTP server shut down successfully")

	// Returning runs the deferred database cleanup.
	return nil
}

/*
Open these URLs in your browser:
- http://localhost:8080/health
- http://localhost:8080/health/db
*/