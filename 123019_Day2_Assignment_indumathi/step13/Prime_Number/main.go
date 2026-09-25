package main

import "fmt"

func main() {
	var n int
	isPrime := true

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	if n < 2 {
		isPrime = false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("Prime number")
	} else {
		fmt.Println("Not a prime number")
	}
}