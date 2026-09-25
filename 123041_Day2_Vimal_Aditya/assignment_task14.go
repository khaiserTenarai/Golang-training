package main

import "fmt"

var globalRole = "Admin"

func main() {

	fmt.Println("\n**************************************")
	fmt.Println("14. Demonstrate variable scope and shadowing.")
	fmt.Println("****************************************")

	name := "Vimal"

	fmt.Println("1. Global scope        :", globalRole)
	fmt.Println("2. Outer function scope:", name)

	if true {
		name := "Aditya"

		fmt.Println("3. Inner block (shadowed):", name)
	}

	// outside the block, the original outer variable is unchanged
	fmt.Println("4. Outer scope again   :", name)
}