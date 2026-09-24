package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var employees []Employee

func main() {
	for {
		fmt.Println("\n===== Employee Management =====")
		fmt.Println("1. Create Employee")
		fmt.Println("2. Read Employees")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			createEmployee()
		case 2:
			readEmployees()
		case 3:
			updateEmployee()
		case 4:
			deleteEmployee()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}

// CREATE
func createEmployee() {
	var employee Employee

	fmt.Print("Enter ID: ")
	fmt.Scan(&employee.ID)

	// Check for duplicate ID
	for _, emp := range employees {
		if emp.ID == employee.ID {
			fmt.Println("Employee ID already exists!")
			return
		}
	}

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Department: ")
	fmt.Scan(&employee.Department)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

	employees = append(employees, employee)

	fmt.Println("Employee created successfully!")
}

// READ
func readEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees found!")
		return
	}

	fmt.Println("\n----- Employee List -----")

	for _, employee := range employees {
		fmt.Printf(
			"ID: %d | Name: %s | Department: %s | Salary: %.2f\n",
			employee.ID,
			employee.Name,
			employee.Department,
			employee.Salary,
		)
	}
}

// UPDATE
func updateEmployee() {
	var id int

	fmt.Print("Enter Employee ID to update: ")
	fmt.Scan(&id)

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {

			fmt.Print("Enter New Name: ")
			fmt.Scan(&employees[i].Name)

			fmt.Print("Enter New Department: ")
			fmt.Scan(&employees[i].Department)

			fmt.Print("Enter New Salary: ")
			fmt.Scan(&employees[i].Salary)

			fmt.Println("Employee updated successfully!")
			return
		}
	}

	fmt.Println("Employee not found!")
}

// DELETE
func deleteEmployee() {
	var id int

	fmt.Print("Enter Employee ID to delete: ")
	fmt.Scan(&id)

	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {

			// Remove employee from slice
			employees = append(
				employees[:i],
				employees[i+1:]...,
			)

			fmt.Println("Employee deleted successfully!")
			return
		}
	}

	fmt.Println("Employee not found!")
}
