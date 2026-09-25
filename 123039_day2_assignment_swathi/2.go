
package main

import "fmt"

func main() {

	// Integer types
	var age int = 25
	var population int64 = 1400000000

	// Floating-point types
	var height float32 = 5.5
	var salary float64 = 50000.75

	// String type
	var name string = "Swathi"

	// Boolean type
	var isEmployee bool = true

	// Byte type
	var letter byte = 'A'

	// Rune type
	var symbol rune = '😊'

	fmt.Println("===== GO PRIMITIVE DATA TYPES =====")

	fmt.Println("int:", age)
	fmt.Println("int64:", population)
	fmt.Println("float32:", height)
	fmt.Println("float64:", salary)
	fmt.Println("string:", name)
	fmt.Println("bool:", isEmployee)
	fmt.Println("byte:", letter)
	fmt.Println("rune:", symbol)
}

