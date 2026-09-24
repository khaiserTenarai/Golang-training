package main

import (
	"fmt"

	"employee-management/employees"
	"employee-management/utils"
)

func main() {
	emp := employee.NewEmployee(101, "Rahul", 50000)

	fmt.Println("Employee Operation")

	utils.PrintEmployee(emp.ID, emp.Name, emp.Salary)
}
