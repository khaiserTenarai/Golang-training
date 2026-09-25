
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	employeeID := uuid.New()

	fmt.Println("===== EMPLOYEE DETAILS =====")
	fmt.Println("Employee ID:", employeeID)
	fmt.Println("Employee Name: Swathi")
}


