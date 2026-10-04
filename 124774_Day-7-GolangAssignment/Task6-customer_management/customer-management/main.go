package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"customer-management/config"
	"customer-management/controller"
	"customer-management/repository"
	"customer-management/service"
	"customer-management/view"
)

func main() {

	fmt.Println(
		"Starting Customer Management System...",
	)

	// --------------------------------------------
	// Load Configuration
	// --------------------------------------------

	cfg := config.Load()

	fmt.Println(
		"Database:",
		cfg.DBName,
	)

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

	fmt.Println(
		"PostgreSQL connected successfully.",
	)

	// ==================================================
	// CUSTOMER DEPENDENCY INJECTION
	// ==================================================

	customerRepository :=
		repository.NewCustomerRepository(db)

	customerService :=
		service.NewCustomerService(
			customerRepository,
		)

	customerView :=
		view.NewCustomerView()

	customerController :=
		controller.NewCustomerController(
			customerService,
			customerView,
		)

	// ==================================================
	// START APPLICATION
	// ==================================================

	customerController.Start()
}
