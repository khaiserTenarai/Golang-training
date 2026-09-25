package main

import "fmt"

func main() {
	var text string = "some one"

	charCount := 0
	wordCount := 0
	vowelCount := 0
	digitCount := 0

	inWord := false

	for i := 0; i < len(text); i++ {
		char := text[i]

		charCount++

		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' ||
			char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' {
			vowelCount++
		}

		if char >= '0' && char <= '9' {
			digitCount++
		}

		if char == ' ' {
			inWord = false
		} else if !inWord {
			inWord = true
			wordCount++
		}
	}

	fmt.Println("Text analyzed:", text)
	fmt.Println("Characters:", charCount)
	fmt.Println("Words:", wordCount)
	fmt.Println("Vowels:", vowelCount)
	fmt.Println("Digits:", digitCount)
}
