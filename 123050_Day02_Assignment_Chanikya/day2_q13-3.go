package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	a := 1

	if n < 2 {
		a = 0
	}

	for i := 2; i < (n/2)+1; i++ {
		if n%i == 0 {
			a = 0
			break
		}
	}

	if a == 1 {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not a prime number")
	}
}
