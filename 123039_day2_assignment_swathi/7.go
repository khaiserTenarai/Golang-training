
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	var text string

	fmt.Println("===== STRING PROCESSING APPLICATION =====")

	fmt.Print("Enter a string: ")
	fmt.Scanln(&text)

	// Count characters
	characters := len([]rune(text))

	// Count words
	words := len(strings.Fields(text))

	// Count vowels and digits
	vowels := 0
	digits := 0

	for _, character := range text {

		// Count digits
		if unicode.IsDigit(character) {
			digits++
		}

		// Count vowels
		if character == 'a' ||
			character == 'e' ||
			character == 'i' ||
			character == 'o' ||
			character == 'u' ||
			character == 'A' ||
			character == 'E' ||
			character == 'I' ||
			character == 'O' ||
			character == 'U' {
			vowels++
		}
	}

	fmt.Println("\n===== RESULT =====")
	fmt.Println("Characters:", characters)
	fmt.Println("Words:", words)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Digits:", digits)
}

