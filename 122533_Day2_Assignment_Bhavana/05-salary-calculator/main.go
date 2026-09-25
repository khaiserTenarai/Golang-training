// 5. Create an employee salary calculator using arithmetic operators.

package main

import "fmt"

func main() {
	basicSalary := 30000.0
	allowance := 5000.0
	deduction := 2000.0

	grossSalary := basicSalary + allowance
	netSalary := grossSalary - deduction
	yearlySalary := netSalary * 12
	monthlyTax := deduction / 12 // just to use division somewhere meaningful

	fmt.Println("Basic Salary:  ", basicSalary)
	fmt.Println("Allowance:     ", allowance)
	fmt.Println("Deduction:     ", deduction)
	fmt.Println("Gross Salary:  ", grossSalary)
	fmt.Println("Net Salary:    ", netSalary)
	fmt.Println("Yearly Salary: ", yearlySalary)
	fmt.Println("Monthly Tax Part:", monthlyTax)
}
