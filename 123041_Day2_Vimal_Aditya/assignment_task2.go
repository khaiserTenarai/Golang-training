package main

import "fmt"


func main(){

	fmt.Println("\n*******************************************")
	fmt.Println("2. Demonstrate all major Go primitive data types. ")
	fmt.Println("*******************************************")

	// BOOLEAN -
	var isActive bool = true

	// INTEGERS
	var num int = 30           
	var smallNum int8 = 127     
	var medNum int16 = 32767  
	var largeNum int32 = 2147483647
	var hugeNum int64 = 9223372036854775807

	// Unsigned integers 
	var count uint = 100
	var byteVal uint8 = 255    

	// FLOATING-POINT NUMBERS
	var pi float32 = 3.1415927
	var price float64 = 99.9999

	// COMPLEX NUMBERS
	var c1 complex64 = 3 + 4i
	var c2 complex128 = 5 + 12i

	// STRING / CHARACTER TYPES
	var message string = "Hello, Go!"
	var letter rune = 'A'

	fmt.Println("=== 1. Booleans ===")
	fmt.Println("Value:", isActive)

	fmt.Println("\n=== 2. Integers ===")
	fmt.Println("int:", num)
	fmt.Println("int8:", smallNum, "| int16:", medNum, "| int32:", largeNum, "| int64:", hugeNum)
	fmt.Println("uint:", count, "| byte (uint8):", byteVal)

	fmt.Println("\n=== 3. Floating-Point ===")
	fmt.Println("float32:", pi, "| float64:", price)

	fmt.Println("\n=== 4. Complex Numbers ===")
	fmt.Println("complex64:", c1, "| complex128:", c2)

	fmt.Println("\n=== 5. Text & Characters ===")
	fmt.Println("string:", message)
	fmt.Println("rune (character):", string(letter), "| Unicode Code Point:", letter)

}