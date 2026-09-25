package main

import "fmt"

func main() {
	var n int
	// fmt.Println(" n value: ",n)

	fmt.Println("Enter a number : ")

	fmt.Scan(&n)

	if IsPrimeNumber(n) {
		fmt.Println(n, " is Prime number ")
	} else {
		fmt.Println(n, " is not a prime number")
	}

}

func IsPrimeNumber(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
