package main

import "fmt"

func main() {

	var num int

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	original := num
	reverse := 0

	for num > 0 {

		digit := num % 10

		reverse = reverse*10 + digit

		num = num / 10
	}

	if original == reverse {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a palindrome")
	}
}
