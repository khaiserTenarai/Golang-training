package main

import (
	"fmt"
	"packagedesign/employee"
	"packagedesign/utils"
)

func main() {
	emp := employee.Employee{
		ID:     101,
		Name:   utils.TitleCase("Sachin"),
		Salary: 45000,
	}

	fmt.Printf("Before raise: %s (%.2f)\n", emp.Name, emp.Salary)
	emp.GiveRaise(10)
	fmt.Printf("After 10%% raise: %s (%.2f)\n", emp.Name, emp.Salary)
}