package main

import "fmt"

func factorial(a int) int {
	if a == 0 {
		return 1
	}
	return a * factorial(a-1)
}
func main() {
	var a int
	var ans int = 1
	fmt.Print("Enter a number: ")
	fmt.Scan(&a)
	fmt.Println("Using Recursion")
	fmt.Printf("%d", factorial(a))

	for i := 1; i <= a; i++ {
		ans *= i
	}
	fmt.Printf("Using Loop = %d\n", ans)
}
