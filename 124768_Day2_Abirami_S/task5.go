package main

import "fmt"

func main() {
	var basicSal, bonus, deduction int
	fmt.Println("Enter Basic Salary: ")
	fmt.Scan(&basicSal)

	fmt.Println("Enter Bonus: ")
	fmt.Scan(&bonus)

	fmt.Println("Enter Deduction: ")
	fmt.Scan(&deduction)

	grossSal := basicSal + bonus
	netSal := grossSal - deduction
	fmt.Println("Basic Salary: ", basicSal)
	fmt.Println("Bonus: ", bonus)
	fmt.Println("Deduction Salary: ", deduction)
	fmt.Println("Gross Salary: ", grossSal)
	fmt.Println("Net Salary: ", netSal)
}
