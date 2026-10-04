package main

import "fmt"

func main() {
	var basicSalary float64
	var bonus float64
	var deduction float64

	fmt.Print("Enter basic salary: ")
	fmt.Scan(&basicSalary)

	fmt.Print("Enter bonus: ")
	fmt.Scan(&bonus)

	fmt.Print("Enter deduction: ")
	fmt.Scan(&deduction)

	grossSalary := basicSalary + bonus
	netSalary := grossSalary - deduction

	fmt.Println("Gross Salary:", grossSalary)
	fmt.Println("Net Salary:", netSalary)
}