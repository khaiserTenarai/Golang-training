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

	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func main() {

	fmt.Println("Recursive:", fibonacciRecursive(6))

	fmt.Println("Iterative:", fibonacciIterative(6))
}