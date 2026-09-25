package main

import "fmt"

func main() {

	n := 12345
	reverse := 0

	for n > 0 {

		digit := n % 10
		reverse = reverse*10 + digit
		n = n / 10
	}

	fmt.Println("Reverse:", reverse)
}
