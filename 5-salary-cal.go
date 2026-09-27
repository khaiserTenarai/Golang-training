package main

import "fmt"

func main() {

	// Employee basic salary
	basicSalary := 30000.0

	// Calculate allowances
	hra := basicSalary * 0.20 // 20% HRA
	da := basicSalary * 0.10  // 10% DA

	// Calculate gross salary
	grossSalary := basicSalary + hra + da

	// Calculate tax
	tax := grossSalary * 0.05 // 5% tax

	// Calculate final salary
	netSalary := grossSalary - tax

	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("HRA:", hra)
	fmt.Println("DA:", da)
	fmt.Println("Gross Salary:", grossSalary)
	fmt.Println("Tax:", tax)
	fmt.Println("Net Salary:", netSalary)
}
