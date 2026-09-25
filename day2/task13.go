package main

import "fmt"

// 1. Factorial
func factorial(n int) uint64 {
	if n < 0 {
		return 0
	}
	var result uint64 = 1
	for i := 1; i <= n; i++ {
		result *= uint64(i)
	}
	return result
}

// 2. Fibonacci Series (displays first n terms)
func fibonacci(terms int) {
	a, b := 0, 1
	fmt.Printf("Fibonacci (%d terms): ", terms)
	for i := 0; i < terms; i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println()
}

// 3. Prime Number Check
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 4. Reverse a Number
func reverseNumber(n int) int {
	reversed := 0
	for n != 0 {
		remainder := n % 10
		reversed = reversed*10 + remainder
		n /= 10
	}
	return reversed
}

// 5. Palindrome Number Check
func isPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	return n == reverseNumber(n)
}

func main() {
	fmt.Println("=== 1. Factorial ===")
	numFact := 5
	fmt.Printf("Factorial of %d = %d\n\n", numFact, factorial(numFact))

	fmt.Println("=== 2. Fibonacci ===")
	fibonacci(7)
	fmt.Println()

	fmt.Println("=== 3. Prime Number ===")
	numPrime := 29
	fmt.Printf("Is %d prime? %t\n\n", numPrime, isPrime(numPrime))

	fmt.Println("=== 4. Reverse Number ===")
	numRev := 12345
	fmt.Printf("Reverse of %d = %d\n\n", numRev, reverseNumber(numRev))

	fmt.Println("=== 5. Palindrome Check ===")
	numPal := 12321
	fmt.Printf("Is %d a palindrome? %t\n", numPal, isPalindrome(numPal))
}