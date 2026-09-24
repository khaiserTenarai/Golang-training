package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Hello Pallavi 123"

	characters := len([]rune(text))
	words := len(strings.Fields(text))

	vowels := 0
	digits := 0

	for _, ch := range text {
		if strings.ContainsRune("aeiouAEIOU", ch) {
			vowels++
		}

		if unicode.IsDigit(ch) {
			digits++
		}
	}

	fmt.Println("Text:", text)
	fmt.Println("Characters:", characters)
	fmt.Println("Words:", words)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Digits:", digits)
}