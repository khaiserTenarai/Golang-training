package main

import "fmt"

func main() {
	var s string
	fmt.Println("Enter a string")

	fmt.Scanf("%s", &s)

	count := 0
	words := 0
	vowels := 0
	digits := 0

	for _, x := range s {
		count++
		if x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' || x == 'A' || x == 'E' || x == 'I' || x == 'O' || x == 'U' {
			vowels++
		}
		if x >= '0' && x <= '9' {
			digits++
		}
		if x == ' ' {
			words++
		}
	}
	fmt.Println("Characters:", count)
	fmt.Println("Words:", words)
	fmt.Println("Vowels:", vowels)
	fmt.Println("Digits:", digits)

}
