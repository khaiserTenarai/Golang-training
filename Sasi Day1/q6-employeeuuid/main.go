package q6employeeuuid
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	employeeID := uuid.New()

	fmt.Println("Employee ID:", employeeID)
}