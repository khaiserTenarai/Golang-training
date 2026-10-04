package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"employee-app/config"
	"employee-app/controller"
	"employee-app/repository"
	"employee-app/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	fmt.Println("Starting Employee Application...")

	cfg := config.Load()

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
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(
			"Database connection failed:",
			err,
		)
	}

	fmt.Println("PostgreSQL connected.")

	// Repository
	employeeRepository :=
		repository.NewEmployeeRepository(db)

	// Repository injected into Service
	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
		)

	// Service injected into Controller
	employeeController :=
		controller.NewEmployeeController(
			employeeService,
		)

	// Start application
	employeeController.Start()
}
