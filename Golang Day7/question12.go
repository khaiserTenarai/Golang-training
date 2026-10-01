package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(10, 5)
	fmt.Println(result)
}