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

	asInt, err := strconv.Atoi(original)
	if err != nil {
		fmt.Println("Could not convert to int:", err)
		return
	}
	fmt.Println("Converted to int:", asInt)

	asFloat := float64(asInt)
	fmt.Println("Converted to float64:", asFloat)

	backToString := strconv.FormatFloat(asFloat, 'f', 2, 64)
	fmt.Println("Back to string:", backToString)
}
