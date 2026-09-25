package main

import "fmt"

func main() {
	x := 99
	fmt.Println("Outer x before block:", x)

	{ // Inner block (creates a new scope)
		x := 10 // Shadows the outer 'x'
		fmt.Println("Inner x inside block :", x)
	}

	// Outer scope resumes
	fmt.Println("Outer x after block :", x)
}
