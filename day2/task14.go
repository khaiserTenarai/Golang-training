package main

import "fmt"

// Package-level scope
var category = "Global Category"

func main() {
	// Function-level scope
	var num = 100
	fmt.Printf("Main scope - num: %d, category: %s\n", num, category)

	{
		num := 500
		category := "Inner Category"

		fmt.Printf("Inner scope (shadowed) - num: %d, category: %s\n", num, category)
	}

	fmt.Printf("Main scope after block - num: %d, category: %s\n", num, category)
}
