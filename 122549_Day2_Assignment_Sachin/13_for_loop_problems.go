package main

import "fmt"

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func fibonacci(n int) []int {
	series := make([]int, 0, n)
	a, b := 0, 1
	for i := 0; i < n; i++ {
		series = append(series, a)
		a, b = b, a+b
	}
	return series
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func reverseNumber(n int) int {
	reversed := 0
	for n != 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}
	return reversed
}

func isPalindrome(n int) bool {
	return n == reverseNumber(n)
}

func main() {
	fmt.Println("Factorial of 5:", factorial(5))
	fmt.Println("First 10 Fibonacci nums:", fibonacci(10))
	fmt.Println("Is 29 prime?:", isPrime(29))
	fmt.Println("Reverse of 1234 :", reverseNumber(1234))
	fmt.Println("Is 121 a palindrome?:", isPalindrome(121))
	fmt.Println("Is 123 a palindrome?:", isPalindrome(123))
}
