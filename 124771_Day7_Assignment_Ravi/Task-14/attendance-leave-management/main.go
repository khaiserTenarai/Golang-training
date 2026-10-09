package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"attendance-leave-management/config"
	"attendance-leave-management/controller"
	"attendance-leave-management/repository"
	"attendance-leave-management/service"
)

func main() {
	fmt.Println("Starting Attendance & Leave Management System...")

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

	attendanceRepository := repository.NewPostgresAttendanceRepository(db)
	attendanceService := service.NewAttendanceService(attendanceRepository)
	attendanceController := controller.NewAttendanceController(attendanceService)

	attendanceController.Start()
}
