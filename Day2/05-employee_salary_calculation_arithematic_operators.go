package main

import "fmt"

func main() {

	basicSalary := 30000.0
	bonus := 5000.0
	deduction := 2000.0

	// Arithmetic operators
	grossSalary := basicSalary + bonus
	netSalary := grossSalary - deduction

	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("Bonus:", bonus)
	fmt.Println("Deduction:", deduction)
	fmt.Println("Gross Salary:", grossSalary)
	fmt.Println("Net Salary:", netSalary)
}
