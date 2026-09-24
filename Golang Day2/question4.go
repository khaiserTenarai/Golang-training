package main

import (
	"fmt"
	"strconv"
)

func main() {
	
	textNumber := "100"

	wholeNumber, _ := strconv.Atoi(textNumber)

	decimalNumber := float64(wholeNumber)

	fmt.Println("Starting Text:   ", textNumber)
	fmt.Println("As a Whole Number:", wholeNumber)
	fmt.Println("As a Decimal:    ", decimalNumber)
}