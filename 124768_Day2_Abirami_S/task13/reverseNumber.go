package main

import "fmt"

func main() {
	var n int
	rev := 0
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)
	for n > 0 {
		last := n % 10
		rev = rev*10 + last
		n = n / 10
	}
	fmt.Println("Reversed Number: ", rev)
}
