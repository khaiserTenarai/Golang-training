package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Go Language was created in 2009 by Google engineers!"

	charCount := 0
	wordCount := len(strings.Fields(text))
	vowelCount := 0
	digitCount := 0

	for _, ch := range text {
		charCount++

		lowerCh := unicode.ToLower(ch)
		if lowerCh == 'a' || lowerCh == 'e' || lowerCh == 'i' || lowerCh == 'o' || lowerCh == 'u' {
			vowelCount++
		}

		if unicode.IsDigit(ch) {
			digitCount++
		}
	}

	fmt.Printf("Input Text: %q\n", text)
	fmt.Printf("Characters: %d\n", charCount)
	fmt.Printf("Words:      %d\n", wordCount)
	fmt.Printf("Vowels:     %d\n", vowelCount)
	fmt.Printf("Digits:     %d\n", digitCount)
}
