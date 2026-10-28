package main

import (
	"context"
	"errors"
	"fmt"
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

// initializeApp creates the configuration, database, and application layers.
func initializeApp() (*config.Config, *gorm.DB, error) {
	// Load normal application settings.
	cfg, err := config.Load()
	if err != nil {
		return nil, nil, fmt.Errorf("load configuration: %w", err)
	}

	// Load sensitive settings separately.
	secrets, err := config.LoadSecrets()
	if err != nil {
		return nil, nil, fmt.Errorf("load secrets: %w", err)
	}

	// Connect to PostgreSQL and run migrations.
	db, err := database.Connect(cfg, secrets)
	if err != nil {
		return nil, nil, fmt.Errorf("initialize database: %w", err)
	}

	return cfg, db, nil
}

// run starts the application and handles shutdown.
func run() error {
	cfg, db, err := initializeApp()
	if err != nil {
		return err
	}

	// Close the database connection pool when the application exits.
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("get SQL database: %w", err)
	}
	defer sqlDB.Close()

	// Initialize repository, service, and controller layers.
	employeeRepo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	// Register routes using the loaded configuration.
	httpServer := server.New(cfg, db, employeeController)

	fmt.Println("Application environment:", cfg.AppEnv)
	fmt.Println("Employee creation enabled:", cfg.EnableEmployeeCreate)
	fmt.Println("Starting server on port:", cfg.AppPort)

	// Listen for Ctrl+C and operating-system termination signals.
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
		fmt.Fprintln(os.Stderr, "Application error:", err)
		os.Exit(1)
	}
}