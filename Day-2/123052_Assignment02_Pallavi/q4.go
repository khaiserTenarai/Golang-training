package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringNumber := "100"

	// String to integer
	number, _ := strconv.Atoi(stringNumber)

	// Integer to float
	floatNumber := float64(number)

	fmt.Println("String:", stringNumber)
	fmt.Println("Integer:", number)
	fmt.Println("Float:", floatNumber)
}