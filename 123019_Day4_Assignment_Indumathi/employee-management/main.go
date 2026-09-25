package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
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

		fmt.Println("Please enter your choice:")

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
				var validationError utility.ValidationError

				if errors.As(err, &validationError) {
					fmt.Println("Validation Error:")
					fmt.Println("Field:", validationError.Field)
					fmt.Println("Message:", validationError.Message)
				} else if errors.Is(err, utility.ErrDuplicateEmployee) {
					fmt.Println("Employee already exists")
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Employee added successfully")
			}

		case 2:
			var id int

			fmt.Println("Enter Employee ID:")
			fmt.Scan(&id)

			employee, err := employeeService.GetEmployee(id)

			if err != nil {
				if errors.Is(err, utility.ErrEmployeeNotFound) {
					fmt.Println("Employee not found")
				} else {
					fmt.Println("Error:", err)
				}
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
				if errors.Is(err, utility.ErrEmployeeNotFound) {
					fmt.Println("Employee not found")
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("Employee deleted successfully")
			}

		case 4:
			var id int

			fmt.Println("Enter Employee ID:")
			fmt.Scan(&id)

			employee, err := employeeService.GetEmployee(id)

			if err != nil {
				if errors.Is(err, utility.ErrEmployeeNotFound) {
					fmt.Println("Employee not found")
				} else {
					fmt.Println("Error:", err)
				}
			} else {
				fmt.Println("\nEmployee Details:")
				fmt.Println("ID:", employee.ID)
				fmt.Println("Name:", employee.Name)
				fmt.Println("Email:", employee.Email)
				fmt.Println("Age:", employee.Age)
				fmt.Println("Salary:", employee.Salary)
			}

		case 5:
			fmt.Println("Exiting Employee Management...")
			return

		default:
			fmt.Println("Wrong choice. Please select 1 to 5.")
		}
	}
}
