package main

import "fmt"

func main() {
	var num int = 12321
	original := num
	reversed := 0

	for temp := num; temp > 0; temp /= 10 {
		remainder := temp % 10
		reversed = (reversed * 10) + remainder
	}

	if original == reversed {
		fmt.Printf("%d is a Palindrome\n", original)
	} else {
		fmt.Printf("%d is NOT a Palindrome\n", original)
	}
}
