package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	num := 5
	fmt.Printf("Factorial of %d is: %d\n", num, factorial(num))
	// Output: Factorial of 5 is: 120
}
