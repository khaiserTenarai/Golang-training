package main

import (
	"fmt"
)

type Employee struct {
	ID     int
	Name   string
	Role   string
	Salary float64
}

var employees []Employee

func main() {
	var choice int

	for {
		fmt.Println("\n============================")
		fmt.Println("  EMPLOYEE MANAGEMENT SYSTEM")
		fmt.Println("============================")
		fmt.Println("1. Add Employee (Create)")
		fmt.Println("2. View All Employees (Read)")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice (1-5): ")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addEmployee()
		case 2:
			viewEmployees()
		case 3:
			updateEmployee()
		case 4:
			deleteEmployee()
		case 5:
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Please enter a number between 1 and 5.")
		}
	}
}

// 1. CREATE
func addEmployee() {
	var emp Employee

	fmt.Print("Enter Employee ID: ")
	fmt.Scanln(&emp.ID)

	fmt.Print("Enter Name: ")
	fmt.Scanln(&emp.Name)

	fmt.Print("Enter Role: ")
	fmt.Scanln(&emp.Role)

	fmt.Print("Enter Salary: ")
	fmt.Scanln(&emp.Salary)

	employees = append(employees, emp)
	fmt.Printf("--> Success: Employee '%s' added!\n", emp.Name)
}

// 2. READ
func viewEmployees() {
	if len(employees) == 0 {
		fmt.Println("\nNo employee records found.")
		return
	}

	fmt.Println("\n--- EMPLOYEE LIST ---")
	for _, emp := range employees {
		fmt.Printf("ID: %d | Name: %-10s | Role: %-10s | Salary: $%.2f\n",
			emp.ID, emp.Name, emp.Role, emp.Salary)
	}
}

// 3. UPDATE
func updateEmployee() {
	var id int
	fmt.Print("Enter Employee ID to update: ")
	fmt.Scanln(&id)

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			fmt.Printf("Updating details for %s (ID: %d):\n", employees[i].Name, id)

			fmt.Print("Enter New Role: ")
			fmt.Scanln(&employees[i].Role)

			fmt.Print("Enter New Salary: ")
			fmt.Scanln(&employees[i].Salary)

			fmt.Println("--> Success: Employee record updated!")
			return
		}
	}

	fmt.Printf("--> Error: Employee with ID %d not found.\n", id)
}

// 4. DELETE
func deleteEmployee() {
	var id int
	fmt.Print("Enter Employee ID to delete: ")
	fmt.Scanln(&id)

	for i, emp := range employees {
		if emp.ID == id {
			// Remove from slice using append trick
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Printf("--> Success: Employee ID %d deleted.\n", id)
			return
		}
	}

	fmt.Printf("--> Error: Employee with ID %d not found.\n", id)
}
