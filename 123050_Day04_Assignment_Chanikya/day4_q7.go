package main

import "fmt"

func fact(a int) int {
	if a == 1 || a == 0 {
		return 1
	}
	return a * fact(a-1)
}
func main() {
	fmt.Print(fact(5))
}
