package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	factorial := 1

	for i := 1; i <= number; i++ {
		factorial = factorial * i
	}

	fmt.Println("Factorial:", factorial)
}