package main

import "fmt"

func main() {

	salaries := []float64{
		30000,
		50000,
		70000,
		40000,
	}

	filterSalary := func(salary float64) bool {
		return salary >= 50000
	}

	fmt.Println("Employees with salary >= 50000:")

	for _, salary := range salaries {

		if filterSalary(salary) {
			fmt.Println(salary)
		}
	}
}