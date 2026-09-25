package main

import "fmt"

func main() {
	var n int
	factorial := 1

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Println("Factorial:", factorial)
}