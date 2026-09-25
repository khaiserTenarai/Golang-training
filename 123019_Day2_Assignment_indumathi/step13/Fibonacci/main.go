package main

import "fmt"

func main() {
	var n int
	first := 0
	second := 1

	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print(first, " ")

		next := first + second
		first = second
		second = next
	}
}