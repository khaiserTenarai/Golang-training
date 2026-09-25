package main

import "fmt"

func fibo_recurr(n int ) int{
	if n<=1{
		return n
	}

	return fibo_recurr(n-1) + fibo_recurr(n-2)
}

func Iterative(n int) int {

	a := 0
	b := 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func main() {

	fmt.Println("Recursive:", fibo_recurr(6))
	fmt.Println("Iterative:", Iterative(6))
}