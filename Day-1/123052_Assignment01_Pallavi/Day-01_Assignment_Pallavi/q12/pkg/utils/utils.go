package utils

import (
	"fmt"

	"q12/internal/employee"
)

func PrintEmployee(emp employee.Employee) {

	fmt.Println("Employee ID:", emp.ID)
	fmt.Println("Employee Name:", emp.Name)
	fmt.Println("Employee Salary:", emp.Salary)
}