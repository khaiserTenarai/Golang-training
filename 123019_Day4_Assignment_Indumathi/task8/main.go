package main

import "fmt"

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}

	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciIterative(n int) int {
	a := 0
	b := 1

	for i := 0; i < n; i++ {
		fmt.Print(a, " ")

		next := a + b
		a = b
		b = next
	}

	return a
}

func main() {
	fmt.Println("Recursive Fibonacci:")
	for i := 0; i < 10; i++ {
		fmt.Print(fibonacciRecursive(i), " ")
	}

	fmt.Println()

	fmt.Println("Iterative Fibonacci:")
	fibonacciIterative(10)
}