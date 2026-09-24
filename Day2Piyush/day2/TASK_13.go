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

	// 2. Fibonacci Series
	terms := 10
	a := 0
	b := 1

	fmt.Print("Fibonacci Series: ")

	for i := 0; i < terms; i++ {
		fmt.Print(a, " ")

		next := a + b
		a = b
		b = next
	}

	fmt.Println()

	// 3. Prime Number Check
	number := 17
	isPrime := true

	if number < 2 {
		isPrime = false
	} else {
		for i := 2; i < number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println(number, "is a Prime Number")
	} else {
		fmt.Println(number, "is not a Prime Number")
	}

	// 4. Reverse a Number
	num := 12345
	reverse := 0
	temp := num

	for temp > 0 {
		digit := temp % 10
		reverse = reverse*10 + digit
		temp = temp / 10
	}

	fmt.Println("Reverse of", num, "=", reverse)

	// 5. Palindrome Check
	palindromeNum := 121
	original := palindromeNum
	reversed := 0

	for palindromeNum > 0 {
		digit := palindromeNum % 10
		reversed = reversed*10 + digit
		palindromeNum = palindromeNum / 10
	}

	if original == reversed {
		fmt.Println(original, "is a Palindrome")
	} else {
		fmt.Println(original, "is not a Palindrome")
	}
}