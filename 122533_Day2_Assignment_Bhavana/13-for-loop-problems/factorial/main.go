// 13a. Factorial
//
// Solve using a for loop: factorial of a number.

package main

import "fmt"

func main() {
	n := 5
	factorial := 1

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Printf("Factorial of %d is %d\n", n, factorial)
}
