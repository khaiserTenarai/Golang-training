package main

import (
	"fmt"

	"Day5Piyush/employee-management/controller"
	"Day5Piyush/employee-management/model"
	"Day5Piyush/employee-management/repository"
	"Day5Piyush/employee-management/service"
)

func main() {

	// Create repository
	employeeRepository := repository.NewInMemoryEmployeeRepository()

	// Create service
	employeeService := service.NewEmployeeService(
		employeeRepository,
	)

	// Create controller
	employeeController := controller.NewEmployeeController(
		employeeService,
	)

	employee := model.Employee{
		ID:       101,
		Name:     "Piyush",
		Email:    "piyush@example.com",
		Age:      24,
		Salary:   50000,
		Phone:    "9876543210",
		Position: "Software Engineer",

		Address: model.Address{
			Street:  "MG Road",
			City:    "Mumbai",
			State:   "Maharashtra",
			ZipCode: "400001",
		},

		Department: model.Department{
			ID:      10,
			Name:    "IT",
			Manager: "Rahul",
		},
	}

	fmt.Println("===== ADD EMPLOYEE =====")

	employeeController.AddEmployee(employee)

	fmt.Println("\n===== GET EMPLOYEE =====")

	employeeController.GetEmployee(101)

	fmt.Println("\n===== ALL EMPLOYEES =====")

	employeeController.GetAllEmployees()

	fmt.Println("\n===== DELETE EMPLOYEE =====")

	employeeController.DeleteEmployee(101)

	fmt.Println("\n===== GET AFTER DELETE =====")

	employeeController.GetEmployee(101)
}