package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"example.com/employee-management/config"
	"example.com/employee-management/controller"
	"example.com/employee-management/repository"
	"example.com/employee-management/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	fmt.Println("Starting Product Inventory System...")

	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := pgxpool.New(ctx, cfg.DatabaseURL())
	if err != nil {
		log.Fatal("Unable to create database pool:", err)
	}
	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal("Unable to connect to PostgreSQL:", err)
	}
	fmt.Println("PostgreSQL connected successfully.")

	productRepository := repository.NewPostgresProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	productController.Start()
}