package employee

import "fmt"

// Employee is an employee record.
type Employee struct {
	ID   int
	Name string
}

// Find returns an employee by ID.
func Find(id int) Employee {
	return Employee{ID: "101", Name: "Asha"}
}
