
package main

import "fmt"

func main() {

	var basicSalary float64
	var bonus float64
	var tax float64

	fmt.Println("===== EMPLOYEE SALARY CALCULATOR =====")

	fmt.Print("Enter basic salary: ")
	fmt.Scan(&basicSalary)

	fmt.Print("Enter bonus: ")
	fmt.Scan(&bonus)

	fmt.Print("Enter tax: ")
	fmt.Scan(&tax)

	// Addition
	grossSalary := basicSalary + bonus

	// Subtraction
	netSalary := grossSalary - tax

	// Multiplication
	yearlySalary := netSalary * 12

	// Division
	monthlyAverage := yearlySalary / 12

	fmt.Println("\n===== SALARY DETAILS =====")
	fmt.Println("Basic Salary:", basicSalary)
	fmt.Println("Bonus:", bonus)
	fmt.Println("Tax:", tax)
	fmt.Println("Gross Salary:", grossSalary)
	fmt.Println("Net Salary:", netSalary)
	fmt.Println("Yearly Salary:", yearlySalary)
	fmt.Println("Monthly Average:", monthlyAverage)
}


