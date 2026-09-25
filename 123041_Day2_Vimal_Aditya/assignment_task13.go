package main

import "fmt"

func main() {

	fmt.Println("\n**************************************")
	fmt.Println("13. Solve five problems using for loops:")
	fmt.Println("****************************************")

	var choice int

	for {
		fmt.Println("\nEnter your Chioce: ")
		fmt.Println("1. Factorial")
		fmt.Println("2. Fibonacci Series")
		fmt.Println("3. Prime Number Check")
		fmt.Println("4. Reverse a Number")
		fmt.Println("5. Palindrome Check")
		fmt.Println("6. Exit")
		fmt.Println("")
		fmt.Print("Choose a problem to solve (1-6): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			solveFactorial()
		case 2:
			solveFibonacci()
		case 3:
			solvePrime()
		case 4:
			solveReverse()
		case 5:
			solvePalindrome()
		case 6:
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Please select 1 to 6.")
		}
	}
}

func solveFactorial() {
	var n int
	fmt.Print("\n[Factorial] Enter a positive integer: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Factorial is not defined for negative numbers.")
		return
	}

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial = factorial * i
	}

	fmt.Printf("Factorial of %d is: %d\n", n, factorial)
}

func solveFibonacci() {
	var n int
	fmt.Print("\n[Fibonacci] Enter number of terms: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Please enter a number greater than 0.")
		return
	}

	a := 0
	b := 1

	fmt.Print("Fibonacci Series: ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", a)
		next := a + b
		a = b
		b = next
	}
	fmt.Println()
}

func solvePrime() {
	var n int
	fmt.Print("\n[Prime Check] Enter a number: ")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Printf("%d is NOT a prime number.\n", n)
		return
	}

	isPrime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Printf("%d IS a prime number.\n", n)
	} else {
		fmt.Printf("%d is NOT a prime number.\n", n)
	}
}

func solveReverse() {
	var n int
	fmt.Print("\n[Reverse Number] Enter a number: ")
	fmt.Scan(&n)

	temp := n
	reversed := 0

	for temp != 0 {
		remainder := temp % 10
		reversed = (reversed * 10) + remainder
		temp = temp / 10
	}

	fmt.Printf("Reversed number of %d is: %d\n", n, reversed)
}

func solvePalindrome() {
	var n int
	fmt.Print("\n[Palindrome Check] Enter a number: ")
	fmt.Scan(&n)

	temp := n
	reversed := 0

	for temp > 0 {
		remainder := temp % 10
		reversed = (reversed * 10) + remainder
		temp = temp / 10
	}

	if n == reversed {
		fmt.Printf("%d IS a palindrome number.\n", n)
	} else {
		fmt.Printf("%d is NOT a palindrome number.\n", n)
	}
}