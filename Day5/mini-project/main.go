package main

import (
	"fmt"

	"example.com/employee-management/controller"
	"example.com/employee-management/repository"
	"example.com/employee-management/service"
)

func main() {
	fmt.Println("Starting Employee Management System...")

	// flow that being used : Repository -> Service -> Controller
	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository)
	employeeController := controller.NewEmployeeController(employeeService)

	employeeController.Start()
}
