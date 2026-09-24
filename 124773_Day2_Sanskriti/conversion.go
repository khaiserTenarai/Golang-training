package main

import (
	"fmt"
	"strconv"
)

func main() {
	// String
	str := "100"

	// String -> Integer
	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Conversion error:", err)
		return
	}

	// Integer -> Float
	floatNum := float64(num)

	fmt.Println("String:", str)
	fmt.Println("Integer:", num)
	fmt.Println("Float:", floatNum)
}