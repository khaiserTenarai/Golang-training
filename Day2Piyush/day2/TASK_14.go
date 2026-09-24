package main

import "fmt"

var global = "Global variable"

func main() {

	// Local variable
	name := "Piyush"

	fmt.Println("Inside main:", name)
	fmt.Println("Global:", global)

	// Inner scope
	if true {
		name := "Rahul"

		fmt.Println("Inside if:", name)
		fmt.Println("Global inside if:", global)
	}

	// Original name is still unchanged
	fmt.Println("Outside if:", name)
}