// 13e. Palindrome
//
// Solve using a for loop: check whether a number reads the same
// forward and backward.

package main

import "fmt"

func main() {
	num := 12321
	reversed := 0
	original := num

	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}

	if reversed == original {
		fmt.Println(original, "is a palindrome")
	} else {
		fmt.Println(original, "is not a palindrome")
	}
}
