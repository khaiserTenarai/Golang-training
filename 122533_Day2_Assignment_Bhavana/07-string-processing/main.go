// 7. Create a string-processing application that counts:
// - Characters
// - Words
// - Vowels
// - Digits

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Go is fun and Go is fast, version 1 rocks!"

	charCount := len(text)
	wordCount := len(strings.Fields(text))

	vowelCount := 0
	digitCount := 0

	for _, ch := range strings.ToLower(text) {
		if strings.ContainsRune("aeiou", ch) {
			vowelCount++
		}
		if unicode.IsDigit(ch) {
			digitCount++
		}
	}

	fmt.Println("Text:", text)
	fmt.Println("Characters:", charCount)
	fmt.Println("Words:", wordCount)
	fmt.Println("Vowels:", vowelCount)
	fmt.Println("Digits:", digitCount)
}
