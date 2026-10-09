package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"ecommerce-order-system/config"
	"ecommerce-order-system/controller"
	"ecommerce-order-system/repository"
	"ecommerce-order-system/service"
)

func main() {
	fmt.Println("Starting E-Commerce Order System...")

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

	ecommerceRepository := repository.NewPostgresEcommerceRepository(db)
	ecommerceService := service.NewEcommerceService(ecommerceRepository)
	ecommerceController := controller.NewEcommerceController(ecommerceService)

	ecommerceController.Start()
}
