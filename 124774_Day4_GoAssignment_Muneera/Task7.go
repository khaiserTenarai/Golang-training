package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	n := 4
	fmt.Printf("Factorial of given number %d is :", n)
	fmt.Println(factorial(n))
}
