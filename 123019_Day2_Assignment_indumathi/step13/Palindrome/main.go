package main

import "fmt"

func main() {
	var n int
	original := 0
	reverse := 0

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	original = n

	for n != 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n /= 10
	}

	if original == reverse {
		fmt.Println("Palindrome number")
	} else {
		fmt.Println("Not a palindrome number")
	}
}