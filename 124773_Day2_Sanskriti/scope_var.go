package main

import "fmt"

var x = "Package Scope" // Package level

func main() {
	x := "Function Scope" // Shadowing package 'x'
	fmt.Println(x)       // Output: Function Scope

	{
		x := "Block Scope" // Shadowing function 'x'
		fmt.Println(x)     // Output: Block Scope
	}

	fmt.Println(x) // Output: Function Scope (original value intact)
}