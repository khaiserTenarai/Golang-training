package main

import "fmt"

// Recursive Fibonacci
func fibonacciRecursive(number int) int {

	if number <= 1 {
		return number
	}

	return fibonacciRecursive(number-1) + fibonacciRecursive(number-2)
}

// Iterative Fibonacci
func fibonacciIterative(number int) int {

	first := 0
	second := 1

	for i := 0; i < number; i++ {
		first, second = second, first+second
	}

	return first
}

func main() {

	number := 6

	fmt.Println("Recursive:", fibonacciRecursive(number))
	fmt.Println("Iterative:", fibonacciIterative(number))
}