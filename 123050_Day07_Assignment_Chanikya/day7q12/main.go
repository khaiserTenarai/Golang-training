package main

import (
	"context"

	"fmt"

	"day7q12/controller"

	"day7q12/repository"

	"day7q12/service"

	"github.com/jackc/pgx/v5"
)

func main() {

	conn, err := pgx.Connect(

		context.Background(),

		"postgres://postgres:admin@localhost:5432/gotraining",
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
