package main

import "fmt"

func main() {
	var basicSalary float64
	var bonus float64
	var deductions float64

	fmt.Print("Enter basic salary: ")
	fmt.Scan(&basicSalary)

	fmt.Print("Enter bonus: ")
	fmt.Scan(&bonus)

	fmt.Print("Enter deductions: ")
	fmt.Scan(&deductions)

	grossSalary := basicSalary + bonus
	netSalary := grossSalary - deductions

	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("Bonus:", bonus)
	fmt.Println("Gross Salary:", grossSalary)
	fmt.Println("Deductions:", deductions)
	fmt.Println("Net Salary:", netSalary)
}