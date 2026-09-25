package main

import "fmt"

func factorial(n int) (fact int) {

	// Base condition
	if n == 1 {
		return 1
	}
	// recusrive function call
	return n * factorial(n-1)
}

func main() {

	var n int
	fmt.Println("Enter n :")

	fmt.Scan(&n)

	fmt.Printf("factorial of  %d is : %d ",n,factorial(n))

	// fmt.Println("factorial of", n, "is:", factorial(n))

}
