package main

import "fmt"

func filterSalaries(salaries []int, condition func(int) bool) []int {
	var result []int
	for _, salary := range salaries {
		if condition(salary) {
			result = append(result, salary)
		}
	}
	return result
}

func main() {
	salaries := []int{45000, 85000, 30000, 120000, 50000}

	highSalaries := filterSalaries(salaries, func(salary int) bool {
		return salary >= 60000
	})

	fmt.Println(highSalaries)
}