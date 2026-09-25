// Add the external package
// go get github.com/google/uuid

package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	employeeID := uuid.New().String()

	fmt.Println("Employee Management")
	fmt.Println("Employee ID:", employeeID)
	fmt.Println("Employee Name: Pooja")
}
