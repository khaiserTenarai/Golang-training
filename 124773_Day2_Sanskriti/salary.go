package main

import "fmt"

func main() {
	basicSalary := 30000.0
	allowance := 5000.0
	bonus := 3000.0
	deduction := 2000.0

	grossSalary := basicSalary + allowance + bonus
	netSalary := grossSalary - deduction

	annualSalary := netSalary * 12

	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("Allowance:", allowance)
	fmt.Println("Bonus:", bonus)
	fmt.Println("Deduction:", deduction)

	fmt.Println("Gross Salary:", grossSalary)
	fmt.Println("Net Salary:", netSalary)
	fmt.Println("Annual Salary:", annualSalary)
}