package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func main() {

	employeeService := service.NewEmployeeService()

	for {
		fmt.Println("\n===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Get Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Exit")

		var choice int

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			addEmployee(employeeService)

		case 2:
			getEmployee(employeeService)

		case 3:
			deleteEmployee(employeeService)

		case 4:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addEmployee(employeeService *service.EmployeeService) {

	var employee model.Employee

	fmt.Println("\n--- Add Employee ---")

	fmt.Print("Enter ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

	err := employeeService.AddEmployee(employee)

	if err != nil {

		var validationErr *utility.ValidationError

		if errors.As(err, &validationErr) {
			fmt.Println("Validation Error:", validationErr)
			return
		}

		if errors.Is(err, service.ErrDuplicateEmployee) {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully!")
}

func getEmployee(employeeService *service.EmployeeService) {

	var id int

	fmt.Println("\n--- Get Employee ---")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	employee, err := employeeService.GetEmployee(id)

	if err != nil {

		if errors.Is(err, service.ErrEmployeeNotFound) {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\n--- Employee Details ---")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
}

func deleteEmployee(employeeService *service.EmployeeService) {

	var id int

	fmt.Println("\n--- Delete Employee ---")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	err := employeeService.DeleteEmployee(id)

	if err != nil {

		if errors.Is(err, service.ErrEmployeeNotFound) {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully!")
}
