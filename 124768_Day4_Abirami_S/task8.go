package main

import "fmt"

func recursiveFibonacci(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return recursiveFibonacci(num-1) + recursiveFibonacci(num-2)
}
func iterativeFibonacci(num int) int {
	a := 0
	b := 1
	for i := 0; i < num; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	var n int
	fmt.Println("Enter the number: ")
	fmt.Scan(&n)
	fmt.Println("Recursive Fibonacci Answer: ", recursiveFibonacci(n))
	fmt.Println("Iterative Fibonacci Answer: ", iterativeFibonacci(n))
}
