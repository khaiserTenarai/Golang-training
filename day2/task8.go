package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	name := "आरव शर्मा" // Aarav Sharma

	
	fmt.Printf("Text: %s\n", name)


	fmt.Printf("Byte Length (len):              %d bytes\n", len(name))

	fmt.Printf("Rune Count (RuneCountInString): %d characters\n", utf8.RuneCountInString(name))

	fmt.Println("\nIterating over bytes:")
	for i := 0; i < len(name); i++ {
		fmt.Printf("%X ", name[i])
	}

	fmt.Println("\n\nIterating over runes (range):")
	for idx, r := range name {
		fmt.Printf("Index %2d: %c (Unicode: %U)\n", idx, r, r)
	}
}