package main

import "fmt"

func updateSalaryValue(salary float64) {
	salary = salary + 5000
	fmt.Println("Salary Value  :", salary)
}

func updateSalaryPointer(salary *float64) {
	*salary = *salary + 5000
	fmt.Println("Salary Pointer:", *salary)
}

func main() {

	fmt.Println("\n****************************************")
	fmt.Println("10. Demonstrate value vs pointer behavior.")
	fmt.Println("******************************************")

	salary := 50000.0

	fmt.Println("Original before calling:", salary)
	updateSalaryValue(salary)
	fmt.Println("Original after calling:", salary)

	fmt.Println("Original before calling:", salary)
	updateSalaryPointer(&salary) 
	fmt.Println("Original after calling:", salary)
}