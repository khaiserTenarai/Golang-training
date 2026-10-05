package main

import (
	"context"
	"example.com/q14-attendance-leave/config"
	"example.com/q14-attendance-leave/controller"
	"example.com/q14-attendance-leave/repository"
	"example.com/q14-attendance-leave/service"
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
	repo := repository.NewPostgresAttendanceRepository(db)
	controller.NewAttendanceController(service.NewAttendanceService(repo)).Start()
}
