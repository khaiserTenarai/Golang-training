package main

import "fmt"

func main() {

	// Employee basic salary
	basicSalary := 30000.0

	// Calculate HRA (20% of basic salary)
	hra := basicSalary * 0.20

	// Calculate DA (10% of basic salary)
	da := basicSalary * 0.10

	// Calculate Gross Salary
	grossSalary := basicSalary + hra + da

	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("HRA:", hra)
	fmt.Println("DA:", da)
	fmt.Println("Gross Salary:", grossSalary)
}