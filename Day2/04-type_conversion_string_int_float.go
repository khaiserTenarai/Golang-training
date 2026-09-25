package main

import (
	"fmt"
	"strconv"
)

func main() {

	// declare a string
	str := "124772"

	// Handling Exception here directly
	// string -->int
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Invalid integer:", err)
		return
	}

	// int --> float

	floatNum := float64(num)

	fmt.Println("String:", str)
	fmt.Println("Integer:", num)
	fmt.Println("Float:", floatNum)
}
