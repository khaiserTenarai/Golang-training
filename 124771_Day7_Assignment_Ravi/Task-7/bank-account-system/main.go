package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"bank-account-system/config"
	"bank-account-system/controller"
	"bank-account-system/repository"
	"bank-account-system/service"
)

func main() {
	fmt.Println("Starting Bank Account System...")

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

	accountRepository := repository.NewPostgresAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)
	accountController := controller.NewAccountController(accountService)

	accountController.Start()
}
