package main

import (
	"fmt"
)

func main() {
	var a bool = true

	var greet string = "Hello"

	var num int = 42
	
	var data byte = 'A' 


	var symbol rune = '#' 

	var pi float64 = 3.14159

	var imaginary complex128 = 5 + 7i

	fmt.Println("Primitive Data Types:")
	fmt.Printf("Boolean: %v \t(Type: %T)\n", a,a)
	fmt.Printf("String : %v \t(Type: %T)\n", greet, greet)
	fmt.Printf("Integer: %v \t\t(Type: %T)\n", num,num)
	fmt.Printf("Byte   : %v \t\t(Type: %T, Character: %c)\n", data,data,data)
	fmt.Printf("Rune   : %v \t(Type: %T, Character: %c)\n", symbol, symbol, symbol)
	fmt.Printf("Float  : %v \t(Type: %T)\n", pi, pi)
	fmt.Printf("Complex: %v \t(Type: %T)\n", imaginary, imaginary)
}