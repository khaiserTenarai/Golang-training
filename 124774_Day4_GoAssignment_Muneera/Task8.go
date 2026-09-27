package main

import "fmt"

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibIterative(n int) {
	a := 0
	b := 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
}
func main() {
	n := 6
	fmt.Println("Recursive Fibonacci")
	for i := 0; i < n; i++ {
		fmt.Print(fibRecursive(i), " ")
	}
	fmt.Println()
	fmt.Println("Iterative Fibonacci")
	fibIterative(n)
}
