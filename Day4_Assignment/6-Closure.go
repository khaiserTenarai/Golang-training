package main

import "fmt"

func salaryIncrement(percent float64) func(float64) float64 {
	return func(salary float64) float64 {
		return salary + salary * percent/100
	}
}


func main() {
	increment := salaryIncrement(10)
	salary := increment(50000)
	fmt.Println(salary)
}