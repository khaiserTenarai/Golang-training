package main

import "fmt"

func main() {
	name := "Shasank Raj"

	// Count bytes
	fmt.Println("Name:", name)
	fmt.Println("Number of bytes:", len(name))

	// Convert string to runes
	runes := []rune(name)

	// Count runes
	fmt.Println("Number of runes:", len(runes))

	// Display each rune
	fmt.Println("Characters:", runes)

}
