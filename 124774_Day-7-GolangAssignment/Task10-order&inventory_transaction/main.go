package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"order-inventory/config"
	"order-inventory/controller"
	"order-inventory/repository"
	"order-inventory/service"
	"order-inventory/view"
)

func main() {

	fmt.Println("Starting Order & Inventory System...")

	cfg := config.Load()

	fmt.Println("Database:", cfg.DBName)

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

	// Repository
	orderRepository :=
		repository.NewOrderRepository(db)

	// Service
	orderService :=
		service.NewOrderService(
			orderRepository,
		)

	// View
	orderView :=
		view.NewOrderView()

	// Controller
	orderController :=
		controller.NewOrderController(
			orderService,
			orderView,
		)

	orderController.Start()
}
