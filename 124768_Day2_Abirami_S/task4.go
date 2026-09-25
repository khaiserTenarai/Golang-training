package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "100"
	num, err := strconv.Atoi(str)
	floatNum := float64(num)
	if err != nil {
		fmt.Println("Conversion Failed")
		return
	}
	fmt.Println("String: ", str)
	fmt.Println("Integer: ", num)
	fmt.Println("Float: ", floatNum)
}
