package main

import "fmt"

func main() {
	
	var basicSalary float64 = 40000.00
	var hra float64 = 15000.00    
	var otherAllowances float64 = 8000.00
	
	var providentFund float64 = 3600.00
	var professionalTax float64 = 200.00

	
	grossSalary := basicSalary + hra + otherAllowances

	totalDeductions := providentFund + professionalTax

	netSalary := grossSalary - totalDeductions

	annualNetSalary := netSalary * 12

	
	fmt.Println("Employee Salary Details")
	fmt.Printf("Basic Salary:      $%.2f\n", basicSalary)
	fmt.Printf("Gross Salary:      $%.2f\n", grossSalary)
	fmt.Printf("Total Deductions:  $%.2f\n", totalDeductions)
	fmt.Printf("Net Monthly Salary:$%.2f\n", netSalary)
	fmt.Printf("Net Annual Salary: $%.2f\n", annualNetSalary)
}