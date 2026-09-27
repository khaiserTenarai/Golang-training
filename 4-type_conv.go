package main

import (
	"fmt"
	"strconv"
)

func main() {

	// String value
	strNumber := "100"

	// Convert string to integer
	intNumber, err := strconv.Atoi(strNumber)

	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}

	// Convert integer to float
	floatNumber := float64(intNumber)

	// Display the values
	fmt.Println("String value:", strNumber)
	fmt.Println("Integer value:", intNumber)
	fmt.Println("Float value:", floatNumber)
}
