package main

import "fmt"

func calculateSalary(salaries ...float64) float64 {

	total := 0.0

	for _, salary := range salaries {
		total = total + salary
	}

	return total
}

func main() {

	total := calculateSalary(
		30000,
		40000,
		50000,
	)

	fmt.Println("Total Salary:", total)
}