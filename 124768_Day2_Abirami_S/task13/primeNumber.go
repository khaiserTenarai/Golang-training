package main

import "fmt"

func main() {
	var n int
	isPrime := true
	fmt.Println("Enter a number: ")
	fmt.Scan(&n)
	if n <= 1 {
		isPrime = false
	} else {
		for i := 2; i <= n/2; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Println("Prime Number")
	} else {
		fmt.Printf("Not a prime number")
	}
}
