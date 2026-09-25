package main

import "fmt"

func increaseSalary(salary *float64, amount float64) {
	*salary = *salary + amount
}

func main() {
	salary := 50000.0

	fmt.Println("Before:", salary)

	increaseSalary(&salary, 5000)

	fmt.Println("After:", salary)
}