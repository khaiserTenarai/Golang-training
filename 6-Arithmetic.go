package main

import "fmt"

func main() {
	// Employee basic salary
	var basicSalary float64 = 30000

	// Allowances
	var a float64 = 5000
	var b float64 = 3000

	// Deductions
	var c float64 = 2000
	var d float64 = 1500

	// Calculate gross salary
	grossSalary := basicSalary + a + b

	// Calculate total deductions
	totalDeduction := c + d

	// Calculate net salary
	netSalary := grossSalary - totalDeduction

	fmt.Println("===== Employee Salary Details =====")
	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("HRA:", a)
	fmt.Println("DA:", b)
	fmt.Println("Gross Salary:", grossSalary)
	fmt.Println("Tax:", c)
	fmt.Println("PF:", d)
	fmt.Println("Total Deduction:", totalDeduction)
	fmt.Println("Net Salary:", netSalary)
}
