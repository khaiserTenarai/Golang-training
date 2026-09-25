package main

import "fmt"

func main() {

	var n int

	fmt.Print("Enter a number: ")
	fmt.Scan(&n)

	fact := 1

	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	fmt.Println("Factorial:", fact)

	a := 0
	b := 1

	fmt.Print("Fibonacci: ")

	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		c := a + b
		a = b
		b = c
	}

	fmt.Println()

	prime := true

	if n < 2 {
		prime = false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prime = false
			break
		}
	}

	if prime {
		fmt.Println("Prime: Yes")
	} else {
		fmt.Println("Prime: No")
	}

	temp := n
	rev := 0

	for temp > 0 {
		d := temp % 10
		rev = rev*10 + d
		temp = temp / 10
	}

	fmt.Println("Reverse:", rev)

	if n == rev {
		fmt.Println("Palindrome: Yes")
	} else {
		fmt.Println("Palindrome: No")
	}
}
