package main

import (
	"fmt"
	"strings"
)

func main() {
	
	text := "Welcome Onboard"

	charCount := 0
	wordCount := 0
	vowelCount := 0
	digitCount := 0

	words := strings.Fields(text)
	wordCount = len(words)

	for _, char := range text {
		
		charCount++ 

		if char >= '0' && char <= '9' {
			digitCount++
		}

		switch char {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowelCount++
		}
	}

	fmt.Println("Total Characters:", charCount)
	fmt.Println("Total Words:     ", wordCount)
	fmt.Println("Total Vowels:    ", vowelCount)
	fmt.Println("Total Digits:    ", digitCount)
}