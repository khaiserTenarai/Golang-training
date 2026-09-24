package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a string: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	characterCount := 0
	wordCount := 0
	vowelCount := 0
	digitCount := 0

	// Count characters, vowels and digits
	for _, char := range input {
		if !unicode.IsSpace(char) {
			characterCount++
		}

		if strings.ContainsRune("aeiouAEIOU", char) {
			vowelCount++
		}

		if unicode.IsDigit(char) {
			digitCount++
		}
	}

	// Count words
	if input != "" {
		wordCount = len(strings.Fields(input))
	}

	fmt.Println("\nResults:")
	fmt.Println("Characters:", characterCount)
	fmt.Println("Words:", wordCount)
	fmt.Println("Vowels:", vowelCount)
	fmt.Println("Digits:", digitCount)
}