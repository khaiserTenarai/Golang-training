package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	characters := 0
	words := 0
	vowels := 0
	digits := 0
	fmt.Println("Enter the sentence: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	isWord := false
	for i := 0; i < len(str); i++ {
		characters++
		if str[i] == 'a' || str[i] == 'e' || str[i] == 'i' || str[i] == 'o' || str[i] == 'u' || str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {
			vowels++
		}
		if str[i] >= '0' && str[i] <= '9' {
			digits++
		}
		if str[i] != ' ' && !isWord {
			words++
			isWord = true
		}
		if str[i] == ' ' {
			isWord = false
		}
	}
	fmt.Println("Characters: ", characters)
	fmt.Println("Words: ", words)
	fmt.Println("Vowels: ", vowels)
	fmt.Println("Digits: ", digits)
}
