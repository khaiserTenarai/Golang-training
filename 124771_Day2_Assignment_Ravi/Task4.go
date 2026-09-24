package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	// String → Integer
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Invalid integer:", err)
		return
	}

	// Integer → Float
	floatNum := float64(num)

	fmt.Println("String:", str)
	fmt.Println("Integer:", num)
	fmt.Println("Float:", floatNum)
}
