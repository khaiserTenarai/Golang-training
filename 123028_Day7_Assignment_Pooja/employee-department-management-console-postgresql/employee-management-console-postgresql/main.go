package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"example.com/employee-management/config"
	"example.com/employee-management/controller"
	"example.com/employee-management/repository"
	"example.com/employee-management/service"
)

func main() {

	fmt.Println("Starting Employee & Department Management System...")

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
		log.Fatal("Unable to create database pool:", err)
	}

	defer db.Close()

	if err := db.Ping(ctx); err != nil {
		log.Fatal(
			"Unable to connect to PostgreSQL:",
			err,
		)
	}

	fmt.Println("PostgreSQL connected successfully.")

	// --------------------------------------------
	// Dependency Injection
	// --------------------------------------------

	// Repositories
	employeeRepository := repository.NewPostgresEmployeeRepository(db)
	departmentRepository := repository.NewPostgresDepartmentRepository(db)

	// Services
	employeeService := service.NewEmployeeService(employeeRepository)
	departmentService := service.NewDepartmentService(departmentRepository)

	// Controller
	appController := controller.NewAppController(employeeService, departmentService)

	// --------------------------------------------
	// Start Console Application
	// --------------------------------------------

	appController.Start()
}