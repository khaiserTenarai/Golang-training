package main

import "fmt"

func GenerateReport(employees []Employee) {
	fmt.Println("Employee Report")
	fmt.Println("----------------")

	for _, employee := range employees {
		fmt.Println(employee.Name, "-", employee.Department)
	}
}