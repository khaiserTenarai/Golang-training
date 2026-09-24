package main

import (
	"fmt"
	"packagedesign/employee"
	"packagedesign/utils"
)

func main() {
	emp := employee.Employee{
		ID:     105,
		Name:   utils.TitleCase("Gokul"),
		Salary: 19000,
	}

	fmt.Printf("Before raise: %s ($%.2f)\n", emp.Name, emp.Salary)
	emp.GiveRaise(25)
	fmt.Printf("After 10%% raise: %s ($%.2f)\n", emp.Name, emp.Salary)
}
