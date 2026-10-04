package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"ecommerce-order/config"
	"ecommerce-order/controller"
	"ecommerce-order/repository"
	"ecommerce-order/service"
	"ecommerce-order/view"
)

func main() {

	fmt.Println("Starting E-Commerce Order System...")

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
