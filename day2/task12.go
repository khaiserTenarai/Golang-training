package main

import (
	"fmt"
	"os"
)

func main() {
	var choice int

	for {
		fmt.Println("\n=== MENU DRIVEN PROGRAM ===")
		fmt.Println("1. Say Hello")
		fmt.Println("2. Display Current Status")
		fmt.Println("3. Calculate Square of a Number")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice (1-4): ")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Hello! Welcome to the Go Menu Program.")
		case 2:
			fmt.Println("System Status: All services running normally.")
		case 3:
			var num int
			fmt.Print("Enter an integer: ")
			fmt.Scanln(&num)
			fmt.Printf("Square of %d is: %d\n", num, num*num)
		case 4:
			fmt.Println("Exiting program. Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice! Please enter a number between 1 and 4.")
		}
	}
}