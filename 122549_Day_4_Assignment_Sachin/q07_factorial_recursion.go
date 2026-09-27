// 7. Implement factorial using recursion.

package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	num := 5
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}
