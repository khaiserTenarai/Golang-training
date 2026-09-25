package main

import "fmt"

func main() {

	n := 10
	a := 0
	b := 1

	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")

		next := a + b
		a = b
		b = next
	}
}
