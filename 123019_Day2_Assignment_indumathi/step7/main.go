package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	characters := len([]rune(input))
	words := len(strings.Fields(input))

	vowels := 0
	digits := 0

	for _, ch := range input {
		if strings.ContainsRune("aeiouAEIOU", ch) {
			vowels++
		}

		if unicode.IsDigit(ch) {
			digits++
		}
	}

	fmt.Println("Characters:", characters)
	fmt.Println("Words:", words)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Digits:", digits)
}