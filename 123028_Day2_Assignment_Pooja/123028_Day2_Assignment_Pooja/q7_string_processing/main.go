
// Q7. Create a string-processing application that counts:

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Golang is fun, learning now"

	charCount := len(text)
	wordCount := len(strings.Fields(text))

	vowelCount := 0
	digitCount := 0

	for _, ch := range text {
		lower := unicode.ToLower(ch)
		if strings.ContainsRune("aeiou", lower) {
			vowelCount++
		}
		if unicode.IsDigit(ch) {
			digitCount++
		}
	}

	fmt.Println("Text :", text)
	fmt.Println("Characters:", charCount)
	fmt.Println("Words :", wordCount)
	fmt.Println("Vowels :", vowelCount)
	fmt.Println("Digits :", digitCount)
}
