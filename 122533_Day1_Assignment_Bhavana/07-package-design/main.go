// 7. Package Design
//
// Create:
// - main
// - employee
// - utils
// Implement a simple employee operation.

package main

import (
	"fmt"

	"packagedesign/employee"
	// "packagedesign/utils"
)

func main() {
	emp := employee.Employee{
		ID:     1,
		Name:   "Bhavana",
		Salary: 35000,
	}

	fmt.Println("Employee Name:", emp.Name)

	fmt.Printf("Employee Salary: Rs. %.2f\n", emp.Salary)
}

// To run commands
// go mod init packagedesign
// go run main.go
