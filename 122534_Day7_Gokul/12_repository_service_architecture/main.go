package main

import (
	"context"
	"example.com/q12-repository-service/config"
	"example.com/q12-repository-service/controller"
	"example.com/q12-repository-service/repository"
	"example.com/q12-repository-service/service"
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
	repo := repository.NewPostgresEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	controller.NewEmployeeController(svc).Start()
}
