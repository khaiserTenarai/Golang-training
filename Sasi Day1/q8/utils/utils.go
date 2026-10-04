package utils

import "fmt"

// PrintEmployee displays employee information.
func PrintEmployee(id int, name string, salary float64) {
	fmt.Println("Employee ID:", id)
	fmt.Println("Employee Name:", name)
	fmt.Println("Employee Salary:", salary)
}