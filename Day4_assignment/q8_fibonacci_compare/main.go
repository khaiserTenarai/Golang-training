
package main

import (
	"fmt"
	"time"
)

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	n := 30

	start := time.Now()
	recResult := fibRecursive(n)
	recDuration := time.Since(start)

	start = time.Now()
	iterResult := fibIterative(n)
	iterDuration := time.Since(start)

	fmt.Printf("Recursive fib(%d) = %d, took %v\n", n, recResult, recDuration)
	fmt.Printf("Iterative fib(%d) = %d, took %v\n", n, iterResult, iterDuration)
}
