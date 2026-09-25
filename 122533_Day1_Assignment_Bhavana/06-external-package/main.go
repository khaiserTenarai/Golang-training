// 6. External Package
//
// Add an external Go dependency.
// Use it in a meaningful example.

// External dependency: github.com/google/uuid
// go get github.com/google/uuid
package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Employee struct {
	ID   string
	Name string
}

func main() {
	emp := Employee{
		ID:   uuid.New().String(),
		Name: "Bhavana",
	}

	fmt.Println("New employee created")
	fmt.Println("ID:  ", emp.ID)
	fmt.Println("Name:", emp.Name)
}
