package main

import "fmt"

func main() {
	basicSalary := 50000.0
	allowance := 5000.0
	bonus := 3000.0
	tax := 10.0
	grossSalary := basicSalary + allowance + bonus

	taxAmount := grossSalary * tax / 100

	// Calculate net salary
	netSalary := grossSalary - taxAmount

	// Average monthly salary component
	average := netSalary / 1

	fmt.Println("Employee Salary Calculator")
	fmt.Printf("Basic Salary : ₹%.2f\n", basicSalary)
	fmt.Printf("Allowance    : ₹%.2f\n", allowance)
	fmt.Printf("Bonus        : ₹%.2f\n", bonus)
	fmt.Printf("Gross Salary : ₹%.2f\n", grossSalary)
	fmt.Printf("Tax (10%%)    : ₹%.2f\n", taxAmount)
	fmt.Printf("Net Salary   : ₹%.2f\n", netSalary)
	fmt.Printf("Average      : ₹%.2f\n", average)
}
