package main

import "fmt"

func totalSalary(salaries ...int) int {
	total := 0

	for _, salary := range salaries {
		total = total + salary
	}

	return total
}

func main() {
	result := totalSalary(30000, 40000, 50000)

	fmt.Println("Total Salary:", result)
}