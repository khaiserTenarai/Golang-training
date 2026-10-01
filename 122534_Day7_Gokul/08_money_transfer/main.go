package main

import (
	"context"
	"example.com/q8-money-transfer/config"
	"example.com/q8-money-transfer/controller"
	"example.com/q8-money-transfer/repository"
	"example.com/q8-money-transfer/service"
	"fmt"
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
	fmt.Println("PostgreSQL connected.")
	repo := repository.NewPostgresAccountRepository(db)
	controller.NewAccountController(service.NewAccountService(repo)).Start()
}
