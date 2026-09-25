package main

import "fmt"

func main() {
	text := "book : Learn python programming language in 90 days!!!"

	chars, words, vowels, digits := 0, 0, 0, 0
	inWord := false

	for _, ch := range text {
		if ch != ' ' {
			chars++
			if !inWord {
				words++
				inWord = true
			}
		} else {
			inWord = false
		}

		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			vowels++
		}

		if ch >= '0' && ch <= '9' {
			digits++
		}
	}

	fmt.Println("Text :", text)
	fmt.Println("Chars:", chars)
	fmt.Println("Words:", words)
	fmt.Println("Vowel:", vowels)
	fmt.Println("Digits:", digits)
}