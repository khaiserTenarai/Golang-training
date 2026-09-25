package main

import "fmt"

var message = "Global variable"

func main() {
	message := "Local variable"

	fmt.Println("Inside main:", message)

	{
		message := "Shadowed variable"
		fmt.Println("Inside block:", message)
	}

	fmt.Println("After block:", message)
}