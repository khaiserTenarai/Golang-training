// Package employee defines the Employee type and helpers.
package employee

import "fmt"

// Employee represents a company employee.
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// New creates a new Employee.
func New(id int, name string, salary float64) Employee {
	return Employee{ID: id, Name: name, Salary: salary}
}

// Describe returns a human-readable summary of the employee.
func (e Employee) Describe() string {
	return fmt.Sprintf("Employee #%d: %s (Salary: %.2f)", e.ID, e.Name, e.Salary)
}
