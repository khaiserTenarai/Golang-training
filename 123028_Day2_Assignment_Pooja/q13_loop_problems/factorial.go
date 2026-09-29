package main

import "fmt"

func main() {

	var num int
	factorial := 1

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		factorial = factorial * i
	}

	fmt.Println("Factorial:", factorial)
}