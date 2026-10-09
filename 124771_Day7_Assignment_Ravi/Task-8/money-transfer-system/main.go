package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"money-transfer-system/config"
	"money-transfer-system/controller"
	"money-transfer-system/repository"
	"money-transfer-system/service"
)

func main() {
	fmt.Println("Starting Money Transfer System...")

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

	transferRepository := repository.NewPostgresTransferRepository(db)
	transferService := service.NewTransferService(transferRepository)
	transferController := controller.NewTransferController(transferService)

	transferController.Start()
}
