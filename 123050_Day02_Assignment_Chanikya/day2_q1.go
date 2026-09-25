package main

import (
	"fmt"
)

var type1 string = "global variable"

func main() {
	type2 := "local variable"
	const pi = 3.14
	fmt.Println("displaying the global variable=", type1)
	fmt.Println("displaying the local variable=", type2)
	fmt.Println("displaying the constant variable=", pi)
}
