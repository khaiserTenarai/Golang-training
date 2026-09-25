package main

import "fmt"

func Fibo(a int) int {
	if a <= 1 {
		return a
	}
	return Fibo(a-1) + Fibo(a-2)
}
func main() {
	fmt.Print(Fibo(10))
}

//it repeatedly calculates the same values. Its time complexity is approximately O(2ⁿ)
