package main

import "fmt"

func main() {
	n := 5
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println("Factorial:", fact)

	// Fibonacci
	a, b := 0, 1
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}

	prime := true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}
	fmt.Println("Prime:", prime)

	reverse := 0
	for n > 0 {
		reverse = reverse*10 + n%10
		n /= 10
	}
	fmt.Println("Reverse:", reverse)

	original, reverse := n, 0
	for n > 0 {
		reverse = reverse*10 + n%10
		n /= 10
	}
	fmt.Println("Palindrome:", original == reverse)
}
