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
)

func main() {
	fmt.Println("Starting Customer Management System...")

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

	customerRepository := repository.NewPostgresCustomerRepository(db)
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	customerController.Start()
}
