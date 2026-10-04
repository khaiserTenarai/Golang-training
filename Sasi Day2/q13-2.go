package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter number of terms: ")
	fmt.Scan(&number)

	first := 0
	second := 1

	for i := 0; i < number; i++ {
		fmt.Print(first, " ")

		next := first + second
		first = second
		second = next
	}
}