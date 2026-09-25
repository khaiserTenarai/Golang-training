// Day 1 - Task 6: External Package
// External dependency: github.com/google/uuid
//go get github.com/google/uuid
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	employeeID := uuid.New()
	fmt.Println("New employee record created.")
	fmt.Println("Employee ID:", employeeID.String())
}
