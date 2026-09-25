package main

import "fmt"

func main() {
	var n int
	reverse := 0

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	for n != 0 {
		digit := n % 10
		reverse = reverse*10 + digit
		n /= 10
	}

	fmt.Println("Reverse:", reverse)
}