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
	// Run the application and report startup or runtime errors.
	if err := run(); err != nil {
		log.Printf("Application stopped: %v", err)
		os.Exit(1)
	}
}

// initializeApp prepares the application dependencies in the correct order.
func initializeApp() (
	*config.Config,
	*gorm.DB,
	*http.Server,
	error,
) {
	// Step 1: Load and validate application configuration.
	log.Println("Step 1: Loading configuration")

	cfg, err := config.Load()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("load configuration: %w", err)
	}

	// Step 2: Load sensitive values separately.
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

	// Step 4: Run database migrations.
	log.Println("Step 4: Running database migrations")

	if err := database.Migrate(db); err != nil {
		// Close the database pool if initialization cannot continue.
		sqlDB, poolErr := db.DB()
		if poolErr == nil {
			_ = sqlDB.Close()
		}

		return nil, nil, nil, fmt.Errorf("run migrations: %w", err)
	}

	// Step 5: Initialize the repository layer.
	log.Println("Step 5: Initializing repository")

	employeeRepository := repository.NewEmployeeRepository(db)

	// Step 6: Initialize the service layer.
	log.Println("Step 6: Initializing service")

	employeeService := service.NewEmployeeService(employeeRepository)

	// Step 7: Initialize the controller layer.
	log.Println("Step 7: Initializing controller")

	employeeController := controller.NewEmployeeController(employeeService)

	// Step 8: Register routes and create the HTTP server.
	log.Println("Step 8: Creating HTTP server")

	httpServer := server.New(cfg, db, employeeController)

	// Return the initialized dependencies to the caller.
	return cfg, db, httpServer, nil
}

// run starts the initialized application and manages its lifecycle.
func run() error {
	// Initialize all dependencies before starting the HTTP server.
	cfg, db, httpServer, err := initializeApp()
	if err != nil {
		return err
	}

	// Get the underlying SQL connection pool for cleanup.
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get database connection pool: %w", err)
	}

	// Close the database pool when run finishes.
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed to close database pool: %v", err)
		} else {
			log.Println("Database connection pool closed")
		}
	}()

	// Listen for Ctrl+C and termination signals.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signalChannel)

	// This context tells the HTTP server when to shut down.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Wait for a shutdown signal in the background.
	go func() {
		receivedSignal := <-signalChannel
		log.Printf("Received shutdown signal: %s", receivedSignal)
		cancel()
	}()

	// Report successful initialization.
	log.Printf(
		"Initialization complete. Starting %s server on port %s",
		cfg.AppEnv,
		cfg.AppPort,
	)

	log.Println("GET  /health")
	log.Println("GET  /health/db")
	log.Println("POST /employees")
	log.Println("GET  /employees/:id")

	// Start serving requests and wait for shutdown.
	if err := server.Start(ctx, httpServer); err != nil {
		return fmt.Errorf("start HTTP server: %w", err)
	}

	log.Println("HTTP server stopped successfully")
	return nil
}


/*
# ============================================================
# TASK 12: STARTUP INITIALIZATION SEQUENCE
# Project: employee-api
# ============================================================

# STEP 1: Go to your project directory.
# Replace this path with your actual project path.
Set-Location "C:\path\to\employee-api"

# STEP 2: Check that Go is installed.
go version

# STEP 3: Format all Go source files.
go fmt ./...

# STEP 4: Download and tidy project dependencies.
go mod tidy

# STEP 5: Compile all packages and check for errors.
go build ./...

# STEP 6: Start the Employee API.
# Make sure PostgreSQL is running and .env has the correct password.
go run .


# STEP 7: Check application health.
Invoke-RestMethod "http://localhost:8080/health"

# STEP 8: Check the database connection.
Invoke-RestMethod "http://localhost:8080/health/db"

# STEP 9: Create an employee.
$body = @{
    name  = "Ganesh"
    email = "ganesh@example.com"
} | ConvertTo-Json

$employee = Invoke-RestMethod `
    -Uri "http://localhost:8080/employees" `
    -Method Post `
    -ContentType "application/json" `
    -Body $body

# Display the created employee and its ID.
$employee

# STEP 10: Retrieve that employee using the returned ID.
Invoke-RestMethod "http://localhost:8080/employees/$($employee.id)"
*/