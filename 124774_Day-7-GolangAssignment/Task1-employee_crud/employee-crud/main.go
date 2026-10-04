package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management-app/config"
	"employee-management-app/controller"
	"employee-management-app/repository"
	"employee-management-app/service"
	"employee-management-app/view"
)

func main() {

	fmt.Println("Starting Employee Management System...")

	// --------------------------------------------
	// Load Configuration
	// --------------------------------------------

	cfg := config.Load()

	fmt.Println("Database:", cfg.DBName)

	// --------------------------------------------
	// Create Database Connection
	// --------------------------------------------

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	db, err := pgxpool.New(
		ctx,
		cfg.DatabaseURL(),
	)

	if err != nil {
		log.Fatal(
			"Unable to create database pool:",
			err,
		)
	}

	defer db.Close()

	// --------------------------------------------
	// Test Database Connection
	// --------------------------------------------

	if err := db.Ping(ctx); err != nil {
		log.Fatal(
			"Unable to connect to PostgreSQL:",
			err,
		)
	}

	fmt.Println("PostgreSQL connected successfully.")

	// --------------------------------------------
	// Dependency Injection
	// --------------------------------------------

	employeeRepository :=
		repository.NewEmployeeRepository(db)

	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
		)

	employeeView :=
		view.NewEmployeeView()

	employeeController :=
		controller.NewEmployeeController(
			employeeService,
			employeeView,
		)

	// --------------------------------------------
	// Start Application
	// --------------------------------------------

	employeeController.Start()
}
