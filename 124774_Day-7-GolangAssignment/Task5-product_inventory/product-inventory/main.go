package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"prodcut-inventory/config"
	"prodcut-inventory/controller"
	"prodcut-inventory/repository"
	"prodcut-inventory/service"
	"prodcut-inventory/view"
)

func main() {

	fmt.Println(
		"Starting Product Inventory System...",
	)

	// Load configuration

	cfg := config.Load()

	// Create database connection

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

	// Test connection

	if err := db.Ping(ctx); err != nil {
		log.Fatal(
			"Unable to connect to PostgreSQL:",
			err,
		)
	}

	fmt.Println(
		"PostgreSQL connected successfully.",
	)

	// Repository

	productRepository :=
		repository.NewProductRepository(db)

	// Service

	productService :=
		service.NewProductService(
			productRepository,
		)

	// View

	productView :=
		view.NewProductView()

	// Controller

	productController :=
		controller.NewProductController(
			productService,
			productView,
		)

	// Start application

	productController.Start()
}
