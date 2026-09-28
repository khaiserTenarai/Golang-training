package main

import "fmt"

func value(num int) int {
	num = 20
	return num
}
func pointer(num *int) int {
	*num = 20
	return *num
}
func main() {
	n := 10
	fmt.Println("Original value: ", n)
	fmt.Println("Value function result: ", value(n))
	fmt.Println("After calling Value function, value of n is: ", n)
	fmt.Println("Pointer function result: ", pointer(&n))
	fmt.Println("After calling Pointer function, value of n is: ", n)
}
