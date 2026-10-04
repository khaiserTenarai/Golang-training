package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"employee-management-app/config"
	"employee-management-app/controller"
	"employee-management-app/repository"
	"employee-management-app/service"
	"employee-management-app/view"
)

func main() {

	fmt.Println("Starting Employee Management System...")

	// --------------------------------------------
	// Load Configuration
	// --------------------------------------------

	cfg := config.Load()

	fmt.Println("Database:", cfg.DBName)

	// --------------------------------------------
	// Create Database Connection
	// --------------------------------------------

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

	// --------------------------------------------
	// Test Database Connection
	// --------------------------------------------

	if err := db.Ping(ctx); err != nil {
		log.Fatal(
			"Unable to connect to PostgreSQL:",
			err,
		)
	}

	fmt.Println("PostgreSQL connected successfully.")

	// ==================================================
	// EMPLOYEE DEPENDENCY INJECTION
	// ==================================================

	employeeRepository :=
		repository.NewEmployeeRepository(db)

	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
		)

	employeeView :=
		view.NewEmployeeView()

	employeeController :=
		controller.NewEmployeeController(
			employeeService,
			employeeView,
		)

	// ==================================================
	// DEPARTMENT DEPENDENCY INJECTION
	// ==================================================

	departmentRepository :=
		repository.NewDepartmentRepository(db)

	departmentService :=
		service.NewDepartmentService(
			departmentRepository,
		)

	departmentView :=
		view.NewDepartmentView()

	departmentController :=
		controller.NewDepartmentController(
			departmentService,
			departmentView,
		)

	// ==================================================
	// MAIN MENU
	// ==================================================

	for {

		fmt.Println()
		fmt.Println("========== MAIN MENU ==========")
		fmt.Println("1. Employee Management")
		fmt.Println("2. Department Management")
		fmt.Println("3. Exit")
		fmt.Println("===============================")

		fmt.Print("Enter your choice: ")

		var choice int

		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		switch choice {

		case 1:

			employeeController.Start()

		case 2:

			departmentController.Start()

		case 3:

			fmt.Println("Thank you. Goodbye!")
			return

		default:

			fmt.Println("Invalid choice.")
		}
	}
}
