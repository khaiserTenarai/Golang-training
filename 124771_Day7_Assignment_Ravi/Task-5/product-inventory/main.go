package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"product-inventory/config"
	"product-inventory/controller"
	"product-inventory/repository"
	"product-inventory/service"
)

func main() {

	fmt.Println(
		"Starting Product Inventory System...",
	)

	cfg := config.Load()

	fmt.Println(
		"Database:",
		cfg.DBName,
	)

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

	productRepository :=
		repository.NewPostgresProductRepository(
			db,
		)

	productService :=
		service.NewProductService(
			db,
			productRepository,
		)

	productController :=
		controller.NewProductController(
			productService,
		)

	// --------------------------------------------
	// Start Console Application
	// --------------------------------------------

	productController.Start()
}
