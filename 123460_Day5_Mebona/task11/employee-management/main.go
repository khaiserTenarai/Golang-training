package main

import (
	"employee-management/controller"
	"employee-management/repository"
	"employee-management/service"
)

func main() {

	employeeRepository := &repository.EmployeeRepositoryImpl{}

	employeeService := service.NewEmployeeService(
		employeeRepository,
	)

	employeeController := controller.NewEmployeeController(
		employeeService,
	)

	employeeController.Start()
}