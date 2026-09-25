// 4. Create a type-conversion program for string -> integer -> float.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	numText := "42"

	// string -> int
	numInt, err := strconv.Atoi(numText)
	if err != nil {
		fmt.Println("could not convert to int:", err)
		return
	}
	fmt.Println("string to int:", numInt)

	// int -> float64
	numFloat := float64(numInt)
	fmt.Println("int to float64:", numFloat)

	// float64 -> string
	backToText := strconv.FormatFloat(numFloat, 'f', 2, 64)
	fmt.Println("float64 to string:", backToText)
}
