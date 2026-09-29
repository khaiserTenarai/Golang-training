// 8. Compare recursive and iterative Fibonacci.

package main

import "fmt"

// recursive 
func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

// iterative 
func fibIterative(n int) int {
	if n <= 1 {
		return n
	}
	first, second := 0, 1
	for i := 2; i <= n; i++ {
		first, second = second, first+second
	}
	return second
}

func main() {
	n := 10

	fmt.Println("Recursive Fibonacci of", n, "is", fibRecursive(n))
	fmt.Println("Iterative Fibonacci of", n, "is", fibIterative(n))

}
