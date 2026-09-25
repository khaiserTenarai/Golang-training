
package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("===== STRING TO INTEGER TO FLOAT =====")

	stringValue := "100"

	fmt.Println("String value:", stringValue)

	// String to Integer using strconv.Atoi()
	intValue, err := strconv.Atoi(stringValue)

	if err != nil {
		fmt.Println("Invalid string. Cannot convert to integer.")
		return
	}

	fmt.Println("Integer value:", intValue)

	// Integer to Float
	floatValue := float64(intValue)

	fmt.Println("Float value:", floatValue)
}

