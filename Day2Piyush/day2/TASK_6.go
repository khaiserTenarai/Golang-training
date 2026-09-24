package main

import "fmt"

func main() {

	a := 10
	b := 20

	// Comparison operators
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// Logical operators
	x := true
	y := false

	fmt.Println("x && y:", x && y)
	fmt.Println("x || y:", x || y)
	fmt.Println("!x:", !x)
}