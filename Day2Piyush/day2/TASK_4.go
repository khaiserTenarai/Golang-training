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
		fmt.Println("Conversion failed")
		return
	}

	fmt.Println("String:", str)
	fmt.Println("Integer:", num)

	// Integer -> Float
	decimal := float64(num)

	fmt.Println("Float:", decimal)
}