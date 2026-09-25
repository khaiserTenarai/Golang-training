package main

import "fmt"

func main() {
	name := "Alice" // Error: name declared but not used
	fmt.Printl(name)
}

// Error 1:
// The variable name was declared but not used.
// Error: declared and not used: name
// Fix: Used the name variable with fmt.Println(name)

// Error 2:
// The function name was written incorrectly as fmt.Printl
// Error: undefined: fmt.printl
// Fix: Changed fmt.printl to fmt.Println
