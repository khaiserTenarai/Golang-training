package main

import "fmt"

func main() {
	var n int
	rev := 0
	org := 0

	fmt.Println("Enter a number: ")
	fmt.Scan(&n)
	org = n

	for n > 0 {
		last := n % 10
		rev = rev*10 + last
		n = n / 10
	}
	if org == rev {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not a Palindrome")
	}
}
