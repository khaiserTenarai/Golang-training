package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&str)

	characters := len(str)
	words := len(strings.Fields(str))
	vowels := 0
	digits := 0

	for _, ch := range str {
		// Count vowels
		if ch == 'a' || ch == 'e' || ch == 'i' ||
			ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' ||
			ch == 'O' || ch == 'U' {
			vowels++
		}

		// Count digits
		if ch >= '0' && ch <= '9' {
			digits++
		}
	}

	fmt.Println("\n--- Result ---")
	fmt.Println("Characters:", characters)
	fmt.Println("Words:", words)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Digits:", digits)
}
