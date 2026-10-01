package main

import (
	"context"
	"example.com/q13-authentication/config"
	"example.com/q13-authentication/controller"
	"example.com/q13-authentication/repository"
	"example.com/q13-authentication/service"
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
	repo := repository.NewPostgresUserRepository(db)
	controller.NewAuthController(service.NewAuthService(repo)).Start()
}
