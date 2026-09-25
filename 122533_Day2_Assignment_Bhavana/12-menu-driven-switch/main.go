// 12. Create a menu-driven program using switch.

package main

import "fmt"

func main() {
	var choice string

	for {
		fmt.Println("\n---- Menu ----")
		fmt.Println("1. Say Hello")
		fmt.Println("2. Show today's tip")
		fmt.Println("3. Show Go mascot name")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		// Read the user's input directly into the 'choice' variable
		fmt.Scanln(&choice)

		switch choice {
		case "1":
			fmt.Println("Hello there!")
		case "2":
			fmt.Println("Tip: always check errors returned by a function.")
		case "3":
			fmt.Println("The Go mascot is called Gopher.")
		case "4":
			fmt.Println("Exiting. Bye!")
			return
		default:
			fmt.Println("Please choose a valid option (1-4).")
		}
	}
}