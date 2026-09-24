// cmd/ Contains the entry points of your application.
// For a CLI application, main.go normally goes here.

package main

import (
	"fmt"

	"employee-app/internal/employee"
	"employee-app/pkg/utils"
)

func main() {
	emp := employee.NewEmployee(
		101,
		"Rahul",
		50000,
	)

	fmt.Println("Employee Management System")
	fmt.Println("--------------------------")

	utils.PrintEmployee(
		emp.ID,
		emp.Name,
		emp.Salary,
	)

	emp.IncreaseSalary(10)

	fmt.Println("\nAfter 10% Salary Increase:")
	utils.PrintEmployee(
		emp.ID,
		emp.Name,
		emp.Salary,
	)

	fmt.Println("\nEmployee Details:")
	fmt.Println(emp.GetDetails())
}
