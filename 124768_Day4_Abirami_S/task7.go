package main

import "fmt"

func factorial(num int) int {
	if num == 0 || num == 1 {
		return 1
	}
	return num * factorial(num-1)
}
func main() {
	var n int
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)
	fmt.Printf("Factorial of %d is: %d", n, factorial(n))
}
