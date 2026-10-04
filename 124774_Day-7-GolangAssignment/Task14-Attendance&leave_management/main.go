package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"attendance-leave/config"
	"attendance-leave/controller"
	"attendance-leave/repository"
	"attendance-leave/service"
	"attendance-leave/view"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	fmt.Println("Starting Attendance & Leave System...")

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

	attendanceRepository :=
		repository.NewAttendanceRepository(db)

	attendanceService :=
		service.NewAttendanceService(
			attendanceRepository,
		)

	attendanceView :=
		view.NewAttendanceView()

	attendanceController :=
		controller.NewAttendanceController(
			attendanceService,
			attendanceView,
		)

	attendanceController.Start()
}
