package main

import "fmt"

func calculateTotalSalary(salaries ...float64) float64 {
	var total float64

	for _, salary := range salaries {
		total += salary
	}

	return total
}

func main() {
	total := calculateTotalSalary(30000, 40000, 50000)

	fmt.Println("Total Salary:", total)
}