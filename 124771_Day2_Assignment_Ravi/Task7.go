package main
import (
	"fmt"
	"strings"
)

func main() {
	var s string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&s)
	characters := len(s)
	words := len(strings.Fields(s))
	vowels := 0
	digits := 0
	for _, ch := range s {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
			vowels++
		}

		if ch >= '0' && ch <= '9' {
			digits++
		}
	}

	fmt.Println("Characters:", characters)
	fmt.Println("Words:", words)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Digits:", digits)
}
