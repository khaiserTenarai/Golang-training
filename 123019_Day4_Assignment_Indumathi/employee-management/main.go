package main

import (
	"fmt"

	"employee-management/model"
	"employee-management/service"
)

func main() {

	employeeService := service.EmployeeService{}

	for {

		fmt.Println("\n**********************")
		fmt.Println("Employee Management")
		fmt.Println("**********************")

		fmt.Println("1. Add Employee")
		fmt.Println("2. Get Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. View Employee")
		fmt.Println("5. Exit")

		fmt.Println("Enter your choice:")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var employee model.Employee

			fmt.Println("Enter Employee ID:")
			fmt.Scan(&employee.ID)

			fmt.Println("Enter Employee Name:")
			fmt.Scan(&employee.Name)

			fmt.Println("Enter Employee Email:")
			fmt.Scan(&employee.Email)

			fmt.Println("Enter Employee Age:")
			fmt.Scan(&employee.Age)

			fmt.Println("Enter Employee Salary:")
			fmt.Scan(&employee.Salary)

			err := employeeService.AddEmployee(employee)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee added successfully")
			}

		case 2:
			var id int

			fmt.Println("Enter Employee ID:")
			fmt.Scan(&id)

			employee, err := employeeService.GetEmployee(id)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee found:")
				fmt.Println("ID:", employee.ID)
				fmt.Println("Name:", employee.Name)
				fmt.Println("Email:", employee.Email)
				fmt.Println("Age:", employee.Age)
				fmt.Println("Salary:", employee.Salary)
			}

		case 3:
			var id int

			fmt.Println("Enter Employee ID:")
			fmt.Scan(&id)

			err := employeeService.DeleteEmployee(id)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee deleted successfully")
			}

		case 4:
			var id int

			fmt.Println("Enter Employee ID:")
			fmt.Scan(&id)

			employeeService.ViewEmployee(id)

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}