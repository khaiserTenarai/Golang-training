package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	result := Add(2, 3)
	// Poorly formatted print statement and unintended Printf verb
	fmt.Printf("Result: %d\n", result)
}
