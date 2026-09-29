// 5. Create an anonymous function for salary filtering.

package main

import "fmt"

func main() {
	salaries := []float64{25000, 52000, 31000, 60000, 18000}

	// this function has no name, it's just stored in a variable
	isHighEarner := func(salary float64) bool {
		return salary > 30000
	}

	var highEarners []float64
	for _, s := range salaries {
		if isHighEarner(s) {
			highEarners = append(highEarners, s)
		}
	}

	fmt.Println("All salaries:", salaries)
	fmt.Println("Salaries above 30000:", highEarners)
}
