package main

import "fmt"

const CompanyName = "Tenarai"

func main() {
	var employeeCount int = 25

	var isHiring = true

	department := "Data Engineering"

	var minSalary, maxSalary float64 = 30000, 90000

	fmt.Println("Company     :", CompanyName)
	fmt.Println("Employees   :", employeeCount)
	fmt.Println("Hiring?     :", isHiring)
	fmt.Println("Department  :", department)
	fmt.Println("Salary Range:", minSalary, "-", maxSalary)
}
