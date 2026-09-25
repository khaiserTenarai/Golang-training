package main

import "fmt"

func main() {

	n := 121
	original := n
	reverse := 0

	for n > 0 {

		digit := n % 10
		reverse = reverse*10 + digit
		n = n / 10
	}

	if original == reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
}
