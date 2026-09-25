package main

import "fmt"

// Global variable
var x = 10

func main() {

	// Local variable
	y := 20

	fmt.Println("Global x:", x)
	fmt.Println("Local y:", y)

	// Start a new block
	{
		// This x shadows the global x
		x := 30

		// This y shadows the y from main()
		y := 40

		fmt.Println("\nInside block:")
		fmt.Println("x:", x)
		fmt.Println("y:", y)
	}

	// The variables inside the block are no longer available
	fmt.Println("\nOutside block:")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
}
