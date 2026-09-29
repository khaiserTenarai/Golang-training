package main

import "fmt"

func main() {

	var n int

	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	first := 0
	second := 1

	for i := 1; i <= n; i++ {

		fmt.Print(first, " ")

		next := first + second
		first = second
		second = next
	}
}
