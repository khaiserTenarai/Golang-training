package main

import "fmt"

func calculateStats(a int, b int) (int, int) {
	sum := a + b
	difference := a - b
	
	return sum, difference
}

func main() {
	
	total, diff := calculateStats(20, 10)

	fmt.Printf("Sum: %d\n", total)
	fmt.Printf("Difference: %d\n", diff)
}