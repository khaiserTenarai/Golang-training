package main

import "fmt"

// Recursive Fibonacci
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

// Iterative Fibonacci
func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	a := 0
	b := 1

	for i := 2; i <= n; i++ {
		next := a + b
		a = b
		b = next
	}

	return b
}

func main() {
	number := 10

	recursiveResult := fibonacciRecursive(number)
	iterativeResult := fibonacciIterative(number)

	fmt.Println("Fibonacci number:", number)
	fmt.Println("Recursive result:", recursiveResult)
	fmt.Println("Iterative result:", iterativeResult)
}
