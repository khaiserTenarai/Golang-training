package main

import "fmt"

func fibo(n int) int {
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
	fmt.Print(fibo(10))
}

//it runs in O(n) time
