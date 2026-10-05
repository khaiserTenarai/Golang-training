package main

import (
	"context"
	"example.com/q10-order-inventory/config"
	"example.com/q10-order-inventory/controller"
	"example.com/q10-order-inventory/repository"
	"example.com/q10-order-inventory/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func main() {
	cfg := config.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, e := pgxpool.New(ctx, cfg.DatabaseURL())
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	if e = db.Ping(ctx); e != nil {
		log.Fatal(e)
	}
	repo := repository.NewPostgresOrderRepository(db)
	controller.NewOrderController(service.NewOrderService(repo)).Start()
}
