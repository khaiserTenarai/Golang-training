package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	fmt.Println("\n********************************************************")
	fmt.Println("7. Create a string-processing application that counts: Characters, Words, Vowels, Digits ")
	fmt.Println("**********************************************************")

	var text string

	fmt.Print("Enter text: ")
	fmt.Scanln(&text)

	runes := []rune(text)

	characters := len(runes)
	words := len(strings.Fields(text))
	vowels := 0
	digits := 0

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		if strings.ContainsRune("aeiouAEIOU", ch) {
			vowels++
		}

		if unicode.IsDigit(ch) {
			digits++
		}
	}

	fmt.Println("\n--- Text Analysis ---")
	fmt.Println("Text:", text)
	fmt.Println("Characters:", characters)
	fmt.Println("Words:", words)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Digits:", digits)
}