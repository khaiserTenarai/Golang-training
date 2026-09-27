package main

import "fmt"

func main() {
	salary := 30000.0
	checkSalary := func(salary float64) bool {
		return salary > 25000
	}
	if checkSalary(salary) {
		fmt.Println("Salary is above 25000")
	} else {
		fmt.Println("Salary is below 25000")
	}

}
