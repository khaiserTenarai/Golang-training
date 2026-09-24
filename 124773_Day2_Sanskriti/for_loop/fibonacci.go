package main

import "fmt"

func printFibonacci(n int) {
	a, b := 0, 1
	fmt.Printf("First %d Fibonacci numbers: ", n)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}

func main() {
	printFibonacci(8)
	// Output: First 8 Fibonacci numbers: 0 1 1 2 3 5 8 13 
}