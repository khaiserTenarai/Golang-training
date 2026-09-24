package main

import (
	"fmt"
)

func main() {
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroPointer *int 

	fmt.Println("--- Zero Values in Go ---")
	fmt.Printf("Zero value for int      : %v\n", zeroInt)
	fmt.Printf("Zero value for float64  : %v\n", zeroFloat)
	fmt.Printf("Zero value for bool     : %v\n", zeroBool)
	
	fmt.Printf("Zero value for string   : %q\n", zeroString) 
	
	fmt.Printf("Zero value for pointer  : %v\n", zeroPointer)
}