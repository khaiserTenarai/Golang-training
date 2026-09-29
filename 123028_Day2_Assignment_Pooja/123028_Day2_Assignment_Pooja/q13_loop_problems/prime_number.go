package main

import "fmt"

func main() {

	var num int
	isPrime := true

	fmt.Print("Enter a number: ")
	fmt.Scan(&num)

	if num <= 1 {
		isPrime = false
	}

	for i := 2; i < num; i++ {

		if num%i == 0 {
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
