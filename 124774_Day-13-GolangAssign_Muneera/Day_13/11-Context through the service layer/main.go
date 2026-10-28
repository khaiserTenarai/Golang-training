package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"employee-api/config"
	"employee-api/controller"
	"employee-api/database"
	"employee-api/repository"
	"employee-api/server"
	"employee-api/service"
)

func main() {
	if err := run(); err != nil {
		log.Printf("Application stopped with error: %v", err)
		os.Exit(1)
	}
}

func run() error {
	// Step 1: Load the single .env file and validate configuration.
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	// Step 2: Load the database password.
	secrets, err := config.LoadSecrets()
	if err != nil {
		return err
	}

	// Step 3: Connect to PostgreSQL.
	db, err := database.Connect(cfg, secrets)
	if err != nil {
		return err
	}

	// Step 4: Get the SQL connection pool for cleanup.
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Failed to close database pool: %v", err)
		} else {
			log.Println("Database connection pool closed")
		}
	}()

	// Step 5: Create the employee table if needed.
	if err := database.Migrate(db); err != nil {
		return err
	}

	// Step 6: Create the repository.
	employeeRepository := repository.NewEmployeeRepository(db)

	// Step 7: Inject the repository into the service.
	employeeService := service.NewEmployeeService(employeeRepository)

	// Step 8: Inject the service into the controller.
	employeeController := controller.NewEmployeeController(employeeService)

	// Step 9: Create the HTTP server.
	httpServer := server.New(cfg, db, employeeController)

	// Step 10: Listen for Ctrl+C and SIGTERM.
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(signalChannel)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		receivedSignal := <-signalChannel
		log.Printf("Received signal: %s", receivedSignal)
		cancel()
	}()

	// Step 11: Start the application.
	log.Printf("Starting %s server on port %s", cfg.AppEnv, cfg.AppPort)
	log.Println("Health endpoint: /health")
	log.Println("Database health endpoint: /health/db")
	log.Println("Create employee: POST /employees")
	log.Println("Get employee: GET /employees/:id")

	if err := server.Start(ctx, httpServer); err != nil {
		return err
	}

	log.Println("HTTP server shut down successfully")
	return nil
}