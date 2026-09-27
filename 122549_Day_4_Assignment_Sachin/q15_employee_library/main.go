// 15. Day 4 Mini Project: Reusable Employee Utility Library

package main

import (
	"fmt"

	"employee-management/service"
)

func main() {
	empService := service.NewEmployeeService()

	emp1, err := empService.AddEmployee("  john DOE  ", 65000)
	if err != nil {
		fmt.Println("Error adding employee:", err)
	} else {
		fmt.Printf("Added: %+v\n", *emp1)
	}

	_, err = empService.AddEmployee("", -100)
	if err != nil {
		fmt.Println("Validation error expected:", err)
	}

	emp, err := empService.GetEmployee(1)
	if err == nil {
		fmt.Printf("Retrieved: %+v\n", *emp)
	}

	fmt.Println("All Employees:", empService.ListEmployees())
}
