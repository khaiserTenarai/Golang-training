package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	number := 5

	result := factorial(number)

	fmt.Println("Factorial of", number, "is:", result)
}
