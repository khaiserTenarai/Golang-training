package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"department-management/config"
	"department-management/controller"
	"department-management/repository"
	"department-management/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	cfg := config.Load()

	db, err := pgxpool.New(
		context.Background(),
		cfg.DatabaseURL(),
	)

	if err != nil {
		fmt.Println(
			"Failed to create database pool:",
			err,
		)
		return
	}

	defer db.Close()

	err = db.Ping(context.Background())

	if err != nil {
		fmt.Println(
			"Failed to connect to PostgreSQL:",
			err,
		)
		return
	}

	fmt.Println(
		"PostgreSQL connected successfully.",
	)

	reader := bufio.NewReader(os.Stdin)

	// Repository layer

	departmentRepository :=
		repository.NewDepartmentRepository(db)

	employeeRepository :=
		repository.NewEmployeeRepository(db)

	// Service layer

	departmentService :=
		service.NewDepartmentService(
			departmentRepository,
		)

	employeeService :=
		service.NewEmployeeService(
			employeeRepository,
			departmentRepository,
		)

	// Controller layer

	departmentController :=
		controller.NewDepartmentController(
			departmentService,
			reader,
		)

	employeeController :=
		controller.NewEmployeeController(
			employeeService,
			reader,
		)

	for {

		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("       DEPARTMENT MANAGEMENT")
		fmt.Println("========================================")
		fmt.Println("1.  Create Department")
		fmt.Println("2.  Get Department")
		fmt.Println("3.  Get All Departments")
		fmt.Println("4.  Update Department")
		fmt.Println("5.  Delete Department")
		fmt.Println("----------------------------------------")
		fmt.Println("6.  Create Employee")
		fmt.Println("7.  Get Employee")
		fmt.Println("8.  Get All Employees")
		fmt.Println("9.  Update Employee")
		fmt.Println("10. Delete Employee")
		fmt.Println("11. Change Employee Department")
		fmt.Println("----------------------------------------")
		fmt.Println("0.  Exit")
		fmt.Println("========================================")

		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		switch input {

		case "1":
			departmentController.Create()

		case "2":
			departmentController.GetByID()

		case "3":
			departmentController.GetAll()

		case "4":
			departmentController.Update()

		case "5":
			departmentController.Delete()

		case "6":
			employeeController.Create()

		case "7":
			employeeController.GetByID()

		case "8":
			employeeController.GetAll()

		case "9":
			employeeController.Update()

		case "10":
			employeeController.Delete()

		case "11":
			employeeController.ChangeDepartment()

		case "0":
			fmt.Println("Application closed.")
			return

		default:
			fmt.Println(
				"Invalid choice. Please try again.",
			)
		}
	}
}
