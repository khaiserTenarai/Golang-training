package main

import (
	"fmt"

	"q12/internal/employee"
	"q12/pkg/utils"
)

func main() {

	emp := employee.Employee{
		ID:     101,
		Name:   "Pallavi",
		Salary: 50000,
	}

	fmt.Println("Employee Details")
	fmt.Println("----------------")

	utils.PrintEmployee(emp)
}