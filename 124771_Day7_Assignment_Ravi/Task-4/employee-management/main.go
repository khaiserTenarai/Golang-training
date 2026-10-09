package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management/config"
	"employee-management/controller"
	"employee-management/repository"
	"employee-management/service"
)

func main() {

	fmt.Println(
		"Starting Salary Management System...",
	)

	// --------------------------------------------
	// Load Configuration
	// --------------------------------------------

	cfg := config.Load()

	fmt.Println("Database:", cfg.DBName)

	// --------------------------------------------
	// Create Context
	// --------------------------------------------

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	// --------------------------------------------
	// PostgreSQL Connection
	// --------------------------------------------

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
	// Test PostgreSQL Connection
	// --------------------------------------------

	if err := db.Ping(ctx); err != nil {

		log.Fatal(
			"Unable to connect to PostgreSQL:",
			err,
		)
	}

	fmt.Println(
		"PostgreSQL connected successfully.",
	)

	// --------------------------------------------
	// Dependency Injection
	// --------------------------------------------

	employeeRepository :=
		repository.NewPostgresSalaryRepository(
			db,
		)

	salaryRepository :=
		repository.NewPostgresSalaryRepository(
			db,
		)

	salaryService :=
		service.NewSalaryService(
			db,
			employeeRepository,
			salaryRepository,
		)

	salaryController :=
		controller.NewSalaryController(
			salaryService,
		)

	// --------------------------------------------
	// Start Console Application
	// --------------------------------------------

	salaryController.Start()
}
