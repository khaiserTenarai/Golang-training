// 8. Compare recursive and iterative Fibonacci.

package main

import "fmt"

func fibonacciRecursive(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}

func fibonacciIterative(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

func main() {
	n := 8
	fmt.Println("Recursive:", fibonacciRecursive(n))
	fmt.Println("Iterative:", fibonacciIterative(n))
}
