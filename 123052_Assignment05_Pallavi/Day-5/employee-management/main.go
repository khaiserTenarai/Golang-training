package main

import (
	"fmt"

	"employee-management/controller"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
)

func main() {

	// Create repository
	employeeRepository := repository.NewEmployeeRepository()

	// Create service
	employeeService := service.NewEmployeeService(
		employeeRepository,
	)

	// Create controller
	employeeController := controller.NewEmployeeController(
		employeeService,
	)

	// Create employee
	employee := model.Employee{
		ID:         101,
		Name:       "Pallavi",
		Email:      "pallavi@gmail.com",
		Age:        22,
		Department: "IT",
		Salary:     29000,
	}

	// Add employee
	employeeController.AddEmployee(employee)

	// View employee
	employeeController.ViewEmployee(101)

	// Add another employee
	employeeController.AddEmployee(
		model.Employee{
			ID:         102,
			Name:       "Sonu",
			Email:      "rahul@gmail.com",
			Age:        25,
			Department: "HR",
			Salary:     35000,
		},
	)

	// View all employees
	employeeController.ViewAllEmployees()

	// Delete employee
	employeeController.DeleteEmployee(102)

	fmt.Println("\nAfter deletion:")

	// View all employees again
	employeeController.ViewAllEmployees()
}