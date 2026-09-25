// 3. Demonstrate zero values for different types.
//
// In Go, a variable that is declared but not given a value automatically
// gets its type's "zero value" instead of being left undefined.

package main

import "fmt"

func main() {
	var i int
	var f float64
	var s string
	var b bool

	fmt.Println("zero value of int:", i)
	fmt.Println("zero value of float64:", f)
	fmt.Printf("zero value of string: %q\n", s)
	fmt.Println("zero value of bool:", b)
}
