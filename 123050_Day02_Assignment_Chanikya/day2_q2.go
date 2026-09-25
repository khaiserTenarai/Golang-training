package main

import "fmt"

func main() {
	var num1 int = 10
	var num2 int8 = 120
	var num3 int16 = 3000
	var num4 int32 = 4000
	var num5 int64 = 5000
	var num6 float32 = 20.99
	var num7 float64 = 64.99
	var num9 complex64 = 6 + 5i
	var num10 complex128 = 5 + 2i
	var myBool bool = true
	var myString string = "Hello, Go!"
	fmt.Printf("%T\n", num1)
	fmt.Printf("%T\n", num2)
	fmt.Printf("%T\n", num3)
	fmt.Printf("%T\n", num4)
	fmt.Printf("%T\n", num5)
	fmt.Printf("%T\n", num6)
	fmt.Printf("%T\n", num7)
	fmt.Printf("%T\n", num9)
	fmt.Printf("%T\n", num10)
	fmt.Printf("%T\n", myBool)
	fmt.Printf("%T\n", myString)
}
