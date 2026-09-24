package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	text := "Go is easy 123"

	// Count characters
	charCount := len([]rune(text))

	// Count words
	wordCount := len(strings.Fields(text))

	// Count vowels and digits
	vowelCount := 0
	digitCount := 0

	for _, ch := range text {

		if strings.ContainsRune("aeiouAEIOU", ch) {
			vowelCount++
		}

		if unicode.IsDigit(ch) {
			digitCount++
		}
	}

	fmt.Println("String:", text)
	fmt.Println("Characters:", charCount)
	fmt.Println("Words:", wordCount)
	fmt.Println("Vowels:", vowelCount)
	fmt.Println("Digits:", digitCount)
}