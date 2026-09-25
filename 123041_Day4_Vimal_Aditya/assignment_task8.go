package main

import "fmt"

func fibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibRecursive(n-1) + fibRecursive(n-2)
}

func fibIterative(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {

	fmt.Println("\n*******************************************")
	fmt.Println("8. Compare recursive and iterative Fibonacci.")
	fmt.Println("*********************************************")

	var n int
	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	fmt.Println("Recursive Result:", fibRecursive(n))
	fmt.Println("Iterative Result :", fibIterative(n))
}