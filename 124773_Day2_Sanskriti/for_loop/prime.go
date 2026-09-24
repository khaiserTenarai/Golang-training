package main

import "fmt"

func main() {
	var num int = 29
	isPrime := true

	if num <= 1 {
		isPrime = false
	} else {
		for i := 2; i <= num/2; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Printf("%d is a Prime Number\n", num)
	} else {
		fmt.Printf("%d is NOT a Prime Number\n", num)
	}
}
