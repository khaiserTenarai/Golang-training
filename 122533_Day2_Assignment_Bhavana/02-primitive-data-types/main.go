// 2. Demonstrate all major Go primitive data types.

package main

import "fmt"

func main() {
	var isActive bool = true
	var age int = 25
	var score float64 = 89.5
	var initial byte = 'A' // byte is an alias for uint8
	var symbol rune = '₹'  // rune is an alias for int32, holds a unicode character
	var name string = "Go Language"

	fmt.Println("bool:  ", isActive)
	fmt.Println("int:   ", age)
	fmt.Println("float64:", score)
	fmt.Println("byte:  ", initial)
	fmt.Println("rune:  ", symbol, string(symbol))
	fmt.Println("string:", name)
}
