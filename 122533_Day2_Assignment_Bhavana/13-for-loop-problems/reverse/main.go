// 13d. Reverse number
//
// Solve using a for loop: reverse the digits of a number.

package main

import "fmt"

func main() {
	num := 12345
	reversed := 0
	original := num

	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}

	fmt.Println("Original number:", original)
	fmt.Println("Reversed number:", reversed)
}
