// Command externaldemo uses the external github.com/google/uuid package
// to assign globally unique IDs to new employees — useful when records are
// created by multiple services and sequential IDs could collide.
package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Employee is an employee with a UUID identifier.
type Employee struct {
	ID   uuid.UUID
	Name string
	Dept string
}

// NewEmployee creates an employee with a random (v4) UUID.
func NewEmployee(name, dept string) Employee {
	return Employee{ID: uuid.New(), Name: name, Dept: dept}
}

func main() {
	staff := []Employee{
		NewEmployee("Asha Rao", "Engineering"),
		NewEmployee("Vikram Iyer", "Finance"),
		NewEmployee("Meera Nair", "HR"),
	}
	for _, e := range staff {
		fmt.Printf("%s  %-12s %s\n", e.ID, e.Name, e.Dept)
	}

	// Validate an ID received from outside (e.g., CLI arg or HTTP request).
	input := "not-a-uuid"
	if _, err := uuid.Parse(input); err != nil {
		fmt.Printf("\nInvalid ID %q rejected: %v\n", input, err)
	}
}
