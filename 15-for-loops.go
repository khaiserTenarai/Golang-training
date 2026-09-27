package main

import "fmt"

func main() {

	// 1. Factorial

	n := 5
	factorial := 1

	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Println("Factorial of", n, "=", factorial)

	// 2. Fibonacci

	a := 0
	b := 1

	fmt.Print("Fibonacci: ")

	for i := 0; i < 10; i++ {
		fmt.Print(a, " ")

		next := a + b
		a = b
		b = next
	}

	fmt.Println()

	// 3. Prime Number

	number := 29
	isPrime := true

	if number < 2 {
		isPrime = false
	}

	for i := 2; i*i <= number; i++ {

		if number%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println(number, "is a Prime Number")
	} else {
		fmt.Println(number, "is not a Prime Number")
	}

	// 4. Reverse Number

	num := 12345
	reverse := 0

	for num > 0 {

		digit := num % 10
		reverse = reverse*10 + digit
		num /= 10
	}

	fmt.Println("Reverse:", reverse)

	// 5. Palindrome

	palindromeNumber := 121
	original := palindromeNumber
	reversed := 0

	for palindromeNumber > 0 {

		digit := palindromeNumber % 10
		reversed = reversed*10 + digit
		palindromeNumber /= 10
	}

	if original == reversed {
		fmt.Println(original, "is a Palindrome")
	} else {
		fmt.Println(original, "is not a Palindrome")
	}
}
