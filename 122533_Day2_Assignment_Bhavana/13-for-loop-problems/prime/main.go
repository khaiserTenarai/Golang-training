// 13c. Prime number
//
// Solve using a for loop: check whether a number is prime.

package main

import "fmt"

func main() {
	num := 29
	isPrime := true

	if num < 2 {
		isPrime = false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(num, "is a prime number")
	} else {
		fmt.Println(num, "is not a prime number")
	}
}
