// 11. Day 5 Mini Project: Refactor Employee Management into
// main, controller, service (with interface), repository (with interface),
// model, methods.

package main

import (
	"fmt"
	"employee-management/controller"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
)

func main() {
	// Initialize layers
	repo := repository.NewInMemoryEmployeeRepository()
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)

	for {
		fmt.Println("\n===============================")
		fmt.Println("  EMPLOYEE MANAGEMENT SYSTEM")
		fmt.Println("===============================")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var name, email string
			var salary float64

			fmt.Print("Enter ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Name (single word): ")
			fmt.Scan(&name)
			fmt.Print("Enter Email: ")
			fmt.Scan(&email)
			fmt.Print("Enter Salary: ")
			fmt.Scan(&salary)

			emp := model.Employee{
				ID:     id,
				Name:   name,
				Email:  email,
				Salary: salary,
			}
			fmt.Println("\n--- Processing ---")
			ctrl.HandleAddEmployee(emp)

		case 2:
			var id int
			fmt.Print("Enter Employee ID to view: ")
			fmt.Scan(&id)
			fmt.Println("\n--- Result ---")
			ctrl.HandleGetEmployee(id)

		case 3:
			var id int
			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&id)
			fmt.Println("\n--- Processing ---")
			ctrl.HandleDeleteEmployee(id)

		case 4:
			fmt.Println("Exiting application. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice, please select a valid option.")
		}
	}
}
