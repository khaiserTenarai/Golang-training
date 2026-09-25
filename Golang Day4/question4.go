package main

import "fmt"

func calculateTotalSalary(base int, bonuses ...int) int {
	total := base
	
	for _, bonus := range bonuses {
		total += bonus
	}
	
	return total
}

func main() {
	salary1 := calculateTotalSalary(50000)
	salary2 := calculateTotalSalary(50000, 2000, 1500)
	salary3 := calculateTotalSalary(60000, 5000, 1000, 500, 200)

	fmt.Println(salary1)
	fmt.Println(salary2)
	fmt.Println(salary3)
}