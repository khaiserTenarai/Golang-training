package main

import (
	"context"
	"errors"
	"fmt"
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

	"gorm.io/gorm"
)

// initializeApp loads configuration and initializes the database.
func initializeApp() (*config.Config, *gorm.DB, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("load configuration: %w", err)
	}

	secrets, err := config.LoadSecrets()
	if err != nil {
		return nil, nil, fmt.Errorf("load secrets: %w", err)
	}

	db, err := database.Connect(cfg, secrets)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize database: %w", err)
	}

	return cfg, db, nil
}

// run starts the application and manages shutdown.
func run() error {
	cfg, db, err := initializeApp()
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get SQL database: %w", err)
	}

	// Close the database pool when the server stops.
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("close database connection: %v", err)
		}
	}()

	// Initialize each application layer.
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	// Register routes using the loaded configuration.
	httpServer := server.New(cfg, db, employeeController)

	log.Printf("Environment: %s", cfg.AppEnv)
	log.Printf("Employee creation enabled: %t", cfg.EnableEmployeeCreate)
	log.Printf("Starting server on port %s", cfg.AppPort)

	// Handle Ctrl+C and SIGTERM.
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	return server.Start(ctx, httpServer)
}

func main() {
	if err := run(); err != nil && !errors.Is(err, context.Canceled) {
		log.Printf("Application stopped with error: %v", err)
		os.Exit(1)
	}
}


/*
Method	Endpoint	Purpose

GET http://localhost:8080/health	Check API status

GET	http://localhost:8080/health/db	
Check database connection


POST	http://localhost:8080/employees	
Create an employee

GET	http://localhost:8080/employees/1	Get employee by ID
*/