package main

import "fmt"

type Employee struct {
	ID    int
	Name  string
	Email string
	Role  string
}

var employees []Employee

func addEmployee() {

	var employee Employee

	fmt.Println("\n===== ADD EMPLOYEE =====")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Employee Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Employee Role: ")
	fmt.Scan(&employee.Role)

	employees = append(employees, employee)

	fmt.Println("Employee added successfully.")
}

func searchEmployee() {

	var id int

	fmt.Println("\n===== SEARCH EMPLOYEE =====")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	for i := 0; i < len(employees); i++ {

		if employees[i].ID == id {
			fmt.Println("Employee ID:", employees[i].ID)
			fmt.Println("Name:", employees[i].Name)
			fmt.Println("Email:", employees[i].Email)
			fmt.Println("Role:", employees[i].Role)
			return
		}
	}

	fmt.Println("Employee not found.")
}

func displayEmployees() {

	fmt.Println("\n===== EMPLOYEE LIST =====")

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for i := 0; i < len(employees); i++ {

		fmt.Println("-------------------------")
		fmt.Println("Employee ID:", employees[i].ID)
		fmt.Println("Name:", employees[i].Name)
		fmt.Println("Email:", employees[i].Email)
		fmt.Println("Role:", employees[i].Role)
	}
}

func deleteEmployee() {

	var id int

	fmt.Println("\n===== DELETE EMPLOYEE =====")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	for i := 0; i < len(employees); i++ {

		if employees[i].ID == id {

			employees = append(employees[:i], employees[i+1:]...)

			fmt.Println("Employee deleted successfully.")
			return
		}
	}

	fmt.Println("Employee not found.")
}

func main() {

	for {

		fmt.Println("\n==============================")
		fmt.Println("   EMPLOYEE MANAGEMENT CLI")
		fmt.Println("==============================")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			addEmployee()

		case 2:
			searchEmployee()

		case 3:
			displayEmployees()

		case 4:
			deleteEmployee()

		case 5:
			fmt.Println("Thank you!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
