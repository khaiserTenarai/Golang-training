package utils

import "fmt"

func PrintEmployee(id int, name string, salary float64) {
	fmt.Println("Employee ID:", id)
	fmt.Println("Employee Name:", name)
	fmt.Printf("Employee Salary: %.2f\n", salary)
}
