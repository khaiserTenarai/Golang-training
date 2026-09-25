package main

import "fmt"

// 3. Zero Values for Different Types

func main() {
	var defaultInt int
	var defaultFloat float64
	var defaultBool bool
	var defaultString string


	fmt.Printf("int:          %d\n", defaultInt)
	fmt.Printf("float64:      %f\n", defaultFloat)
	fmt.Printf("bool:         %t\n", defaultBool)
	fmt.Printf("string:       %q (empty string)\n", defaultString)

}
