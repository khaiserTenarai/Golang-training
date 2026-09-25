package main

import (
	"fmt"
)

func CalculateSalary(a ...int) int {
	t := 0
	for _, b := range a {
		t += b
	}
	return t
}
func main() {
	total := CalculateSalary(30000, 5000, 3000, 2000)

	fmt.Println("Total salary:", total)
}
