package main

import "fmt"

func main() {
	// Integer
	var age int = 22

	// Floating-point
	var salary float64 = 45000.75

	// Boolean
	var isEmployee bool = true

	// String
	var name string = "Sanskriti"

	// Complex number
	var complexNumber complex128 = 10 + 5i

	// Byte - alias for uint8
	var letter byte = 'A'

	// Rune - alias for int32
	var symbol rune = '₹'

	fmt.Println("Integer:", age)
	fmt.Println("Float:", salary)
	fmt.Println("Boolean:", isEmployee)
	fmt.Println("String:", name)
	fmt.Println("Complex:", complexNumber)
	fmt.Println("Byte:", letter)
	fmt.Println("Rune:", symbol)
}