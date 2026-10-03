package main

import (
	"context"

	"fmt"

	"assignment_task12/controller"

	"assignment_task12/repository"

	"assignment_task12/service"

	"github.com/jackc/pgx/v5"
)

func main() {

	conn, err := pgx.Connect(

		context.Background(),

		"postgres://postgres:Info%40131@localhost:5432/go_training",
	)

	if err != nil {

		panic(err)

	}

	defer conn.Close(context.Background())

	// Repository Layer

	employeeRepository := repository.NewEmployeeRepository(conn)

	// Service Layer

	employeeService := service.NewEmployeeService(
		employeeRepository,
	)

	// Controller Layer

	employeeController := controller.NewEmployeeController(
		employeeService,
	)

	fmt.Println("===== CREATE EMPLOYEE =====")

	employeeController.CreateEmployee()

	fmt.Println("\n===== GET EMPLOYEES =====")

	employeeController.GetEmployees()

}
