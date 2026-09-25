package main

import "fmt"

func main() {

	var hotel string = "Asha Tiffins"
	var choice int

	fmt.Println("Welcome to", hotel)
	fmt.Println("You can select food from the menu below")

	for {

		fmt.Println(">>>>> South Indian Food Menu <<<<<")
		fmt.Println("1. South Indian Meals")
		fmt.Println("2. Bisibelebath")
		fmt.Println("3. Dosa")
		fmt.Println("4. Idli")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			fmt.Println("You selected South Indian Meals")

		case 2:
			fmt.Println("You selected Bisibelebath")

		case 3:
			fmt.Println("You selected Dosa")

		case 4:
			fmt.Println("You selected Idli")

		case 5:
			fmt.Println("Thank you! Visit again.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
