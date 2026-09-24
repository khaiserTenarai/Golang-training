package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("\n1. Enter a number for Factorial: ")
	fmt.Scan(&n)
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Printf("Factorial of %d is: %d\n", n, fact)

	
	var terms int
	fmt.Print("\n2. Enter number of Fibonacci terms to generate: ")
	fmt.Scan(&terms)
	a, b := 0, 1
	fmt.Print("Fibonacci sequence: ")
	for i := 0; i < terms; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()

	var num int
	fmt.Print("\n3. Enter a number to check if it is Prime: ")
	fmt.Scan(&num)
	isPrime := true
	if num < 2 {
		isPrime = false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Printf("%d is a prime number? %v\n", num, isPrime)

	var val int
	fmt.Print("\n4. Enter a number to reverse: ")
	fmt.Scan(&val)
	origVal := val
	rev := 0
	for val > 0 {
		rev = (rev * 10) + (val % 10)
		val /= 10
	}
	fmt.Printf("Reverse of %d is: %d\n", origVal, rev)

	var orig int
	fmt.Print("\n5. Enter a number to check for Palindrome: ")
	fmt.Scan(&orig)
	temp := orig
	r := 0
	for temp > 0 {
		r = (r * 10) + (temp % 10)
		temp /= 10
	}
	
	if orig == r {
		fmt.Printf("%d is a palindrome\n", orig)
	} else {
		fmt.Printf("%d is not a palindrome\n", orig)
	}
}