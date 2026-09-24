package main

import (
	"fmt"

	"q8/employee"
	"q8/utils"
)

func main() {
	emp := employee.CreateEmployee(101, "Pallavi", 50000)

	fmt.Println("Employee Details")
	utils.PrintEmployee(emp.ID, emp.Name, emp.Salary)
}