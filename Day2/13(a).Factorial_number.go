package main

import "fmt"

func main() {

	n := 5
	factorial := 1

	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}

	fmt.Println("Factorial:", factorial)
}
