package main

import "fmt"

func modifiedValue(num *int) int {
	*num = 20
	return *num
}
func main() {
	n := 10
	p := &n
	fmt.Println("Value of n: ", n)
	fmt.Println("Address of n: ", p)
	fmt.Println("Value of n using *: ", *p)
	fmt.Println("Modified value: ", modifiedValue(p))
}
