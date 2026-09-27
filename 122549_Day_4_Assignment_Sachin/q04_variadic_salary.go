// 4. Create a variadic salary calculation function.

package main

import "fmt"

func calculateTotalSalary(baseSalary float64, bonuses ...float64) float64 {
	total := baseSalary
	for _, bonus := range bonuses {
		total += bonus
	}
	return total
}

func main() {
	s1 := calculateTotalSalary(50000)
	s2 := calculateTotalSalary(50000, 2000, 1500, 500)

	fmt.Println("Base only:", s1)
	fmt.Println("With bonuses:", s2)
}
