package main

import "fmt"

func main() {
	salaries := []float64{
		25000,
		35000,
		45000,
		55000,
		65000,
	}

	filterSalary := func(salary float64) bool {
		return salary >= 50000
	}

	fmt.Println("Salaries greater than or equal to 50000:")

	for _, salary := range salaries {
		if filterSalary(salary) {
			fmt.Println(salary)
		}
	}
}