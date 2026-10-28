package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"employee-api/config"
	"employee-api/controller"
	"employee-api/database"
	"employee-api/repository"
	"employee-api/server"
	"employee-api/service"

	"gorm.io/gorm"
)

func main() {
	// Run the application and report any startup or runtime error.
	if err := run(); err != nil {
		log.Printf("Application stopped with error: %v", err)
		os.Exit(1)
	}
}

// initializeApp creates the dependencies required by the application.
func initializeApp() (
	*config.Config,
	*gorm.DB,
	*http.Server,
	error,
) {
	// Step 1: Load and validate configuration.
	log.Println("Step 1: Loading configuration")

	cfg, err := config.Load()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load configuration: %w", err)
	}

	// Step 2: Load secrets separately.
	log.Println("Step 2: Loading secrets")

	secrets, err := config.LoadSecrets()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load secrets: %w", err)
	}

	// Step 3: Connect to PostgreSQL.
	log.Println("Step 3: Connecting to PostgreSQL")

	db, err := database.Connect(cfg, secrets)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("connect database: %w", err)
	}

	// Close the database if a later initialization step fails.
	startupComplete := false
	defer func() {
		if !startupComplete {
			sqlDB, poolErr := db.DB()
			if poolErr == nil {
				_ = sqlDB.Close()
			}
		}
	}()

	// Step 4: Run database migrations.
	log.Println("Step 4: Running database migrations")

	if err := database.Migrate(db); err != nil {
		return nil, nil, nil, fmt.Errorf("run migrations: %w", err)
	}

	// Step 5: Initialize the repository.
	log.Println("Step 5: Initializing repository")

	employeeRepository := repository.NewEmployeeRepository(db)

	// Step 6: Initialize the service.
	log.Println("Step 6: Initializing service")

	employeeService := service.NewEmployeeService(employeeRepository)

	// Step 7: Initialize the controller.
	log.Println("Step 7: Initializing controller")

	employeeController := controller.NewEmployeeController(employeeService)

	// Step 8: Register routes and create the HTTP server.
	log.Println("Step 8: Creating HTTP server")

	httpServer := server.New(cfg, db, employeeController)

	startupComplete = true
	return cfg, db, httpServer, nil
}

// run manages the application lifecycle.
func run() error {
	// Initialize everything before the server starts.
	cfg, db, httpServer, err := initializeApp()
	if err != nil {
		return err
	}

	// Get the database connection pool for cleanup.
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get database connection pool: %w", err)
	}

	// Listen for Ctrl+C and SIGTERM.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signalChannel)

	// The context tells the HTTP server when to shut down.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Wait for a shutdown signal in the background.
	go func() {
		receivedSignal := <-signalChannel
		log.Printf("Received shutdown signal: %s", receivedSignal)
		cancel()
	}()

	// Log startup details without exposing secrets.
	log.Printf(
		"Application ready: environment=%s port=%s",
		cfg.AppEnv,
		cfg.AppPort,
	)

	// Start the HTTP server and wait for it to stop.
	serverErr := server.Start(ctx, httpServer)

	// Close PostgreSQL only after the HTTP server has stopped.
	log.Println("Closing database connection pool")

	if err := sqlDB.Close(); err != nil {
		if serverErr != nil {
			return fmt.Errorf(
				"server error: %v; database cleanup error: %w",
				serverErr,
				err,
			)
		}
		return fmt.Errorf("close database connection pool: %w", err)
	}

	log.Println("Database connection pool closed")

	if serverErr != nil {
		return serverErr
	}

	log.Println("Application shutdown completed successfully")
	return nil
}