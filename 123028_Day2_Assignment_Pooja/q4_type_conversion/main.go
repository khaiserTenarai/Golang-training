/*
Day 2, Q4. Create a type-conversion program for string -> integer -> float.
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	original := "125"
	fmt.Println("Original string:", original)

	n, err := strconv.Atoi(original)
	if err != nil {
		fmt.Println("Could not convert to int:", err)
		return
	}
	fmt.Println("Converted to int:", n)

	asFloat := float64(n)
	fmt.Println("Converted to float64:", asFloat)

	backToString := strconv.FormatFloat(asFloat, 'f', 2, 64)
	fmt.Println("Back to string:", backToString)
}
