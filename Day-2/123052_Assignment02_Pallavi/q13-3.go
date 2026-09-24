package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	isPrime := true

	if number < 2 {
		isPrime = false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
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