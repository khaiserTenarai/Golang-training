package main

import "fmt"

func main() {
	var integer int
	var decimal float64
	var flag bool
	var text string
	var character rune
	var data []int
	var pointer *int

	fmt.Println("Integer zero value:", integer)
	fmt.Println("Float zero value:", decimal)
	fmt.Println("Boolean zero value:", flag)
	fmt.Println("String zero value:", text)
	fmt.Println("Rune zero value:", character)
	fmt.Println("Slice zero value:", data)
	fmt.Println("Pointer zero value:", pointer)
}