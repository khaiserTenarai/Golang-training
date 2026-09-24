package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

func main() {

	// Create employees
	employees := []Employee{
		{ID: 1, Name: "Piyush", Salary: 25000},
		{ID: 2, Name: "Prachi", Salary: 50000},
		{ID: 3, Name: "yash", Salary: 35000},
	}

	// READ - Display employees
	fmt.Println("Employees:")
	for _, emp := range employees {
		fmt.Println(emp.ID, emp.Name, emp.Salary)
	}

	// CREATE - Add a new employee
	newEmployee := Employee{
		ID:     4,
		Name:   "Priya",
		Salary: 45000,
	}

	employees = append(employees, newEmployee)

	fmt.Println("\nAfter Adding Employee:")
	for _, emp := range employees {
		fmt.Println(emp.ID, emp.Name, emp.Salary)
	}

	// UPDATE - Change salary of employee ID 2
	for i := range employees {
		if employees[i].ID == 2 {
			employees[i].Salary = 50000
		}
	}

	fmt.Println("\nAfter Updating Employee:")
	for _, emp := range employees {
		fmt.Println(emp.ID, emp.Name, emp.Salary)
	}

	// DELETE - Delete employee ID 3
	for i := range employees {
		if employees[i].ID == 3 {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}

	fmt.Println("\nAfter Deleting Employee:")
	for _, emp := range employees {
		fmt.Println(emp.ID, emp.Name, emp.Salary)
	}
}