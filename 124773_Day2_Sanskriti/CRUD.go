package main

import "fmt"

// Employee defines the structure for our data
type Employee struct {
	ID     int
	Name   string
	Role   string
	Salary float64
}

// Slice to store employees globally in memory
var employees []Employee

// 1. CREATE: Add a new employee to the slice
func createEmployee(id int, name string, role string, salary float64) {
	newEmp := Employee{
		ID:     id,
		Name:   name,
		Role:   role,
		Salary: salary,
	}
	employees = append(employees, newEmp)
	fmt.Printf("Added: %s (ID: %d)\n", name, id)
}

// 2. READ: Display all employees
func readEmployees() {
	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println("\n--- Employee List ---")
	for _, emp := range employees {
		fmt.Printf("ID: %d | Name: %-10s | Role: %-10s | Salary: $%.2f\n",
			emp.ID, emp.Name, emp.Role, emp.Salary)
	}
	fmt.Println("---------------------")
}

// 3. UPDATE: Modify an existing employee's role and salary by ID
func updateEmployee(id int, newRole string, newSalary float64) {
	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees[i].Role = newRole
			employees[i].Salary = newSalary
			fmt.Printf("Updated ID %d: New Role = %s, New Salary = $%.2f\n", id, newRole, newSalary)
			return
		}
	}
	fmt.Printf("Employee with ID %d not found.\n", id)
}

// 4. DELETE: Remove an employee by ID from the slice
func deleteEmployee(id int) {
	for i, emp := range employees {
		if emp.ID == id {
			// Remove element using slice append trick
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Printf("Deleted Employee ID: %d\n", id)
			return
		}
	}
	fmt.Printf("Employee with ID %d not found.\n", id)
}

func main() {
	// --- Demonstration ---

	// Create
	fmt.Println("=== Creating Employees ===")
	createEmployee(101, "Alice", "Developer", 75000)
	createEmployee(102, "Bob", "Designer", 65000)
	createEmployee(103, "Charlie", "Manager", 85000)

	// Read
	readEmployees()

	// Update
	fmt.Println("\n=== Updating Employee ===")
	updateEmployee(102, "Sr. Designer", 72000)
	readEmployees()

	// Delete
	fmt.Println("\n=== Deleting Employee ===")
	deleteEmployee(101)
	readEmployees()
}