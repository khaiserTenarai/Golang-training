package main

import (
	"context"
	"example.com/q9-ecommerce-order/config"
	"example.com/q9-ecommerce-order/controller"
	"example.com/q9-ecommerce-order/repository"
	"example.com/q9-ecommerce-order/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func main() {
	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := pgxpool.New(ctx, cfg.DatabaseURL())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	repo := repository.NewPostgresOrderRepository(db)
	controller.NewOrderController(service.NewOrderService(repo)).Start()
}
