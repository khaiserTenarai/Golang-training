
// Day 4, Q4. Create a variadic salary calculation function.

package main

import "fmt"

func totalSalary(salaries ...float64) float64 {
	total := 0.0
	for _, s := range salaries {
		total += s
	}
	return total
}

func main() {
	single := totalSalary(50000)
	fmt.Println("Single salary total:", single)

	teamTotal := totalSalary(50000, 45000, 60000, 55000)
	fmt.Println("Team salary total  :", teamTotal)

	figures := []float64{40000, 42000, 39000}
	fmt.Println("From slice  :", totalSalary(figures...))
}
