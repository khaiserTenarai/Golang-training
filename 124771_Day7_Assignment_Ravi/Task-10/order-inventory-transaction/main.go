package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"order-inventory-transaction/config"
	"order-inventory-transaction/controller"
	"order-inventory-transaction/repository"
	"order-inventory-transaction/service"
)

func main() {
	fmt.Println("Starting Order & Inventory Transaction System...")

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

	orderRepository := repository.NewPostgresOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	orderController.Start()
}
