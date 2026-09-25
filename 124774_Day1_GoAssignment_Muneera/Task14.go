package main

import "fmt"

func main() {
	name := "Muneera"

	fmt.Println("Name:", name)
}

// Error 1:
// Problem: undefined: name

// Cause:
// The variable name was not declared before using it.

// Incorrect:
// name = "Muneera"

// Fix:
// name := "Muneera"

// The := operator declares the variable and assigns the value "Muneera".

// Error 2:
// Problem: Syntax error in argument list.

// Cause:
// A comma was missing between "Name:" and name.

// Incorrect:
// fmt.Println("Name:" name)

// Fix:
// fmt.Println("Name:", name)

// The comma separates the two arguments passed to fmt.Println().
