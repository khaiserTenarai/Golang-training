package main

import "fmt"

func recursive(n int) int {
	if n <= 1 {
		return n
	}
	return recursive(n-1) + recursive(n-2)
}

func iterative(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	return a
}

func main() {
	fmt.Print("Recursive: ")
	for i := 0; i < 7; i++ {
		fmt.Print(recursive(i), " ")
	}

	fmt.Print("\nIterative: ")
	iterative(7)
}