package main

import "fmt"

func main() {
	var baseSalary,bonus,deductions int

	fmt.Println("Enter base salary:")
    fmt.Scanln(&baseSalary)
	fmt.Println("Enter bonus amount:")
    fmt.Scanln(&bonus)
	fmt.Println("Enter salary deduction:")
    fmt.Scanln(&deductions)

	grossSalary := baseSalary + bonus

	tax := (grossSalary*15)/100

	netSalary := grossSalary-tax-deductions

	fmt.Println("Gross Salary :", grossSalary)
    fmt.Println("Tax Amount   :", tax)
    fmt.Println("Deductions   :", deductions)
    fmt.Println("Net Salary   :", netSalary)


}