package main

import "fmt"

func main() {
	var num int = 12345
	original := num
	reversed := 0

	for num > 0 {
		remainder := num % 10
		reversed = (reversed * 10) + remainder
		num /= 10
	}

	fmt.Printf("Original: %d | Reversed: %d\n", original, reversed)
}
