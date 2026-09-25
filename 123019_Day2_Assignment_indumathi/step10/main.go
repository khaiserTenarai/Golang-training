package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var employees []Employee

// Create
func createEmployee(emp Employee) {
	employees = append(employees, emp)
	fmt.Println("Employee created successfully")
}

// Read
func readEmployees() {
	for _, emp := range employees {
		fmt.Println("ID:", emp.ID)
		fmt.Println("Name:", emp.Name)
		fmt.Println("Department:", emp.Department)
		fmt.Println("Salary:", emp.Salary)
		fmt.Println()
	}
}

// Update
func updateEmployee(id int, name string, department string, salary float64) {
	for i := range employees {
		if employees[i].ID == id {
			employees[i].Name = name
			employees[i].Department = department
			employees[i].Salary = salary

			fmt.Println("Employee updated successfully")
			return
		}
	}

	fmt.Println("Employee not found")
}

// Delete
func deleteEmployee(id int) {
	for i := range employees {
		if employees[i].ID == id {
			employees = append(employees[:i], employees[i+1:]...)

			fmt.Println("Employee deleted successfully")
			return
		}
	}

	fmt.Println("Employee not found")
}

func main() {

	// CREATE
	createEmployee(Employee{
		ID:         101,
		Name:       "Indumathi",
		Department: "Data Engineering",
		Salary:     50000,
	})

	createEmployee(Employee{
		ID:         102,
		Name:       "Arun",
		Department: "IT",
		Salary:     45000,
	})

	// READ
	fmt.Println("Employee Details:")
	readEmployees()

	// UPDATE
	updateEmployee(101, "Indumathi A", "Data Engineering", 55000)

	// READ
	fmt.Println("After Update:")
	readEmployees()

	// DELETE
	deleteEmployee(102)

	// READ
	fmt.Println("After Delete:")
	readEmployees()
}