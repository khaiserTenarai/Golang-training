package main

import "fmt"

func main() {
	salaries := []float64{25000, 45000, 30000, 60000, 40000}

	filterSalary := func(salary float64) bool {
		return salary >= 40000
	}

	for _, salary := range salaries {
		if filterSalary(salary) {
			fmt.Println(salary)
		}
	}
}
