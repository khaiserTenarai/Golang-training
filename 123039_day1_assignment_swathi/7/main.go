// Creating main.go

package main

import (
	"employeeapp/employee"
	"employeeapp/utils"
)

func main() {

	utils.PrintMessage("===== EMPLOYEE MANAGEMENT =====")

	employee.AddEmployee("Swathi", 101)
}