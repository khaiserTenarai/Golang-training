package main

import "fmt"

// Employee structure
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Add a new employee
func addEmployee(employees []Employee) []Employee {
	var id int
	var name string
	var salary float64

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&name)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&salary)

	// Add employee to slice
	employees = append(employees, Employee{id, name, salary})

	fmt.Println("Employee added successfully.")

	return employees
}

// Display all employees
func viewEmployees(employees []Employee) {
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println("\n--- Employee List ---")

	// Loop through employees
	for _, emp := range employees {
		fmt.Println("ID:", emp.ID)
		fmt.Println("Name:", emp.Name)
		fmt.Println("Salary:", emp.Salary)
		fmt.Println()
	}
}

// Search employee using map
func searchEmployee(employeeMap map[int]Employee) {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	// Check if employee exists
	employee, found := employeeMap[id]

	if found {
		fmt.Println("\nEmployee Found")
		fmt.Println("ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Salary:", employee.Salary)
	} else {
		fmt.Println("Employee not found.")
	}
}

// Delete employee from slice
func deleteEmployee(employees []Employee) []Employee {
	var id int

	fmt.Print("Enter Employee ID to delete: ")
	fmt.Scan(&id)

	for i, emp := range employees {

		if emp.ID == id {

			// Remove employee from slice
			employees = append(employees[:i], employees[i+1:]...)

			fmt.Println("Employee deleted successfully.")

			return employees
		}
	}

	fmt.Println("Employee not found.")

	return employees
}

func main() {

	// Slice to store employees
	employees := []Employee{}

	// Map for quick employee lookup
	employeeMap := make(map[int]Employee)

	var choice int

	// Menu loop
	for {

		fmt.Println("\n==============================")
		fmt.Println("   EMPLOYEE MANAGEMENT CLI")
		fmt.Println("==============================")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Search Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Println("==============================")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			// Add employee to slice
			employees = addEmployee(employees)

			// Update map with the latest employee
			employee := employees[len(employees)-1]
			employeeMap[employee.ID] = employee

		case 2:
			// Display employees
			viewEmployees(employees)

		case 3:
			// Search employee using map
			searchEmployee(employeeMap)

		case 4:
			// Delete employee from slice
			var id int
			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&id)

			found := false

			// Find employee in slice
			for i, emp := range employees {
				if emp.ID == id {

					// Delete from slice
					employees = append(employees[:i], employees[i+1:]...)

					// Delete from map
					delete(employeeMap, id)

					fmt.Println("Employee deleted successfully.")

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 5:
			fmt.Println("Thank you! Program ended.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
