
package main

import "fmt"

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

func fibonacci(number int) {
	first := 0
	second := 1

	fmt.Print("Fibonacci: ")

	for i := 0; i < number; i++ {
		fmt.Print(first, " ")

		next := first + second
		first = second
		second = next
	}

	fmt.Println()
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func reverseNumber(number int) int {
	reversed := 0

	for number > 0 {
		digit := number % 10
		reversed = reversed*10 + digit
		number = number / 10
	}

	return reversed
}

func isPalindrome(number int) bool {
	original := number
	reversed := reverseNumber(number)

	if original == reversed {
		return true
	}

	return false
}

func main() {
	fmt.Println("===== FIVE FOR LOOP PROBLEMS =====")

	number := 5

	fmt.Println()
	fmt.Println("1. FACTORIAL")
	fmt.Println("Number:", number)
	fmt.Println("Factorial:", factorial(number))

	fmt.Println()
	fmt.Println("2. FIBONACCI")
	fibonacci(10)

	primeNumber := 17

	fmt.Println()
	fmt.Println("3. PRIME NUMBER")
	fmt.Println("Number:", primeNumber)

	if isPrime(primeNumber) {
		fmt.Println("17 is a prime number")
	} else {
		fmt.Println("17 is not a prime number")
	}

	fmt.Println()
	fmt.Println("4. REVERSE NUMBER")
	fmt.Println("Original number:", 12345)
	fmt.Println("Reversed number:", reverseNumber(12345))

	palindromeNumber := 121

	fmt.Println()
	fmt.Println("5. PALINDROME")
	fmt.Println("Number:", palindromeNumber)

	if isPalindrome(palindromeNumber) {
		fmt.Println("121 is a palindrome")
	} else {
		fmt.Println("121 is not a palindrome")
	}
}

