package main

import "fmt"

func main() {

	fmt.Println("\n******************************************************************")
	fmt.Println("5. Create an employee salary calculator using arithmetic operators. ")
	fmt.Println("********************************************************************")

	var Salary float64
	var bonus float64
	var tax float64

	fmt.Print("Enter Basic Salary: ")
	fmt.Scan(&Salary)

	fmt.Print("Enter Bonus: ")
	fmt.Scan(&bonus)

	fmt.Print("Enter Tax Rate (%): ")
	fmt.Scan(&tax)

	// Arithmetic Calculations
	grossSalary := Salary + bonus
	taxDeduction := grossSalary * (tax / 100)
	netSalary := grossSalary - taxDeduction

	fmt.Println("\n--- Salary Summary ---")

	fmt.Println("Basic Salary  :", Salary)
	fmt.Println("Bonus         :", bonus)

	fmt.Println("----------------------")

	fmt.Println("Gross Salary  :", grossSalary)
	fmt.Println("Tax Deduction :", taxDeduction)
	fmt.Println("Net Salary    :", netSalary)

}