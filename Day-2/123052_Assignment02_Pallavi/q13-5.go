package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	original := number
	reverse := 0

	for number > 0 {
		digit := number % 10
		reverse = reverse*10 + digit
		number = number / 10
	}

	if original == reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}