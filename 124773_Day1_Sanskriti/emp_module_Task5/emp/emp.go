package employee

import "fmt"

// Employee represents a worker in the system.
type Employee struct {
	ID   int
	Name string
	Role string
}

// DisplayInfo prints the details of the employee.
func (e Employee) DisplayInfo() {
	fmt.Printf("ID: %d | Name: %s | Role: %s\n", e.ID, e.Name, e.Role)
}

// New returns a new Employee instance with default role.
func New(id int, name string) Employee {
	return Employee{
		ID:   id,
		Name: name,
		Role: "Software Engineer",
	}
}
