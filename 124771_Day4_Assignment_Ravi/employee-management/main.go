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
	// Dummy Employees

	employees := []model.Employee{
		{
			ID:     1,
			Name:   "Ravi Ranjan",
			Email:  "raviranjan@gmail.com",
			Age:    30,
			Salary: 75000,
		},
		{
			ID:     2,
			Name:   "Amit",
			Email:  "amit@gmail.com",
			Age:    28,
			Salary: 65000,
		},
	}

	for _, employee := range employees {
		err := employeeService.AddEmployee(employee)

		if err != nil {
			fmt.Println("Error:", err)
		}
	}

	fmt.Println("================================")
	fmt.Println(" Employee Management System")
	fmt.Println("================================")
	fmt.Println("2 dummy employees added.")

	// Menu

	for {
		fmt.Println("\nChoose an operation:")
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
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// Add Employee

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
		// errors.As() for custom ValidationError
		var validationErr *utility.ValidationError

		if errors.As(err, &validationErr) {
			fmt.Println("\nValidation Error")
			fmt.Println("Field:", validationErr.Field)
			fmt.Println("Message:", validationErr.Message)
			return
		}

		if errors.Is(err, service.ErrDuplicateEmployee) {
			fmt.Println("\nError: Employee already exists.")
			return
		}

		fmt.Println("\nError:", err)
		return
	}

	fmt.Println("\nEmployee added successfully.")
}

// Get Employee

func getEmployee(employeeService *service.EmployeeService) {
	var id int

	fmt.Println("\n--- Get Employee ---")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	employee, err := employeeService.GetEmployee(id)

	if err != nil {
		if errors.Is(err, service.ErrEmployeeNotFound) {
			fmt.Println("\nEmployee not found.")
			return
		}

		fmt.Println("\nError:", err)
		return
	}

	fmt.Println("\nEmployee Details")
	fmt.Println("----------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Printf("Salary: %.2f\n", employee.Salary)
}

// Delete Employee

func deleteEmployee(employeeService *service.EmployeeService) {
	var id int

	fmt.Println("\n--- Delete Employee ---")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	err := employeeService.DeleteEmployee(id)

	if err != nil {
		if errors.Is(err, service.ErrEmployeeNotFound) {
			fmt.Println("\nEmployee not found.")
			return
		}

		fmt.Println("\nError:", err)
		return
	}

	fmt.Println("\nEmployee deleted successfully.")
}
