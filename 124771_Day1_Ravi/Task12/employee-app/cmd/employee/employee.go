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

	utils.PrintEmployee(
		emp.ID,
		emp.Name,
		emp.Salary,
	)

	fmt.Println()
	fmt.Println("Employee Details:")
	fmt.Println(emp.GetDetails())
}
