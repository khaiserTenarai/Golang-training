
package main

import "fmt"

func main() {

	// Variables
	var employeeName string = "Swathi"
	var employeeAge int = 25
	var employeeSalary float64 = 50000.50

	// Constants
	const companyName string = "Tenarai Technologies"
	const country string = "India"

	fmt.Println("===== EMPLOYEE DETAILS =====")
	fmt.Println("Employee Name:", employeeName)
	fmt.Println("Employee Age:", employeeAge)
	fmt.Println("Employee Salary:", employeeSalary)

	fmt.Println("\n===== CONSTANT DETAILS =====")
	fmt.Println("Company Name:", companyName)
	fmt.Println("Country:", country)

	// Changing a variable
	employeeAge = 26

	fmt.Println("\nAfter changing employee age:")
	fmt.Println("Employee Age:", employeeAge)
}

