package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	strv := "20"
	fmt.Printf("Original String:\t Value: %s,\t Type: %T\n", strv, strv)

	intv, err := strconv.Atoi(strv)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
		return
	}
	fmt.Printf("String Converted to Int:\t Value: %d,\t Type: %T\n", intv, intv)

	floatv := float64(intv)
	fmt.Printf("Int Converted to Float:\t Value: %.2f,\t Type: %T\n", floatv, floatv)
}