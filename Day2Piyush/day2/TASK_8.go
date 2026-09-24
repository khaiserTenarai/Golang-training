package main

import "fmt"

func main() {

	text := "Go"

	// Length in bytes
	fmt.Println("String:", text)
	fmt.Println("Number of bytes:", len(text))

	// Length in runes (characters)
	runes := []rune(text)
	fmt.Println("Number of runes:", len(runes))

	// Print bytes
	fmt.Println("Bytes:", []byte(text))

	// Print runes
	fmt.Println("Runes:", runes)

	// Print each rune as a character
	fmt.Println("Characters:")

	for _, r := range text {
		fmt.Printf("%c ", r)
	}

	fmt.Println()
}