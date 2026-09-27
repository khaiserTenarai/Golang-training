package day1assignment
package main

import "fmt"

// Global variable
var message = "Global Message"

func main() {

	// Local variable
	age := 22

	fmt.Println("Global message:", message)
	fmt.Println("Age:", age)

	// Inner block
	if true {

		// This creates a NEW variable.
		// It shadows the outer age variable.
		age := 25

		fmt.Println("Inside block:", age)
	}

	// Outer age is unchanged
	fmt.Println("Outside block:", age)

	// Shadowing the global variable
	message := "Local Message"

	fmt.Println("Inside main:", message)
}