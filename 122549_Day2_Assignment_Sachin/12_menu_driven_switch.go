package main

import "fmt"

func main() {
	var choice int

	for {
		fmt.Println("1. Say Hello")
		fmt.Println("2. See Task")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("hope you're having a good day.")
		case 2:
			fmt.Println("Today's task: finish the Go assignment.")
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Please choose opt 1, 2, or 3.")
		}
	}
}