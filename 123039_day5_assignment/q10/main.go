package main

import (
	"fmt"

	"example.com/employee_management/controller"
	"example.com/employee_management/repository"
	"example.com/employee_management/service"
)

func main() {

	fmt.Println("Starting Employee Management System")

	employeeRepository := repository.NewEmployeeRepository()

	employeeService := service.NewEmployeeService(employeeRepository)

	employeeController := controller.NewEmployeeController(employeeService)

	employeeController.Start()
}
