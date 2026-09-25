package main

import "fmt"

func changeValue(n int) {
	n = 100
}

func changePointer(n *int) {
	*n = 100
}

func main() {
	x := 10

	fmt.Println("Original:", x)

	changeValue(x)
	fmt.Println("After value:", x)

	changePointer(&x)
	fmt.Println("After pointer:", x)
}
