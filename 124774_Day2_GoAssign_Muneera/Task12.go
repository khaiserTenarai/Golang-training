package main

import "fmt"

func main() {
	var choice int
	var continueChoice string

	for {
		fmt.Println("\n********* MAIN MENU ***********")
		fmt.Println("1. Display Employee Name")
		fmt.Println("2. Display Employee Age")
		fmt.Println("3. Display Employee Designation")
		fmt.Println("4. Exit")

		fmt.Print("Enter your Choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			fmt.Println("Name: Ayesha")

		case 2:
			fmt.Println("Age: 21")

		case 3:
			fmt.Println("Designation: Developer")

		case 4:
			fmt.Println("Exiting the program")
			return

		default:
			fmt.Println("Invalid Choice")
		}

		fmt.Print("Do you want to continue? (yes/no): ")
		fmt.Scanf("%s", &continueChoice)

		if continueChoice != "yes" {
			fmt.Println("Program ended")
			break
		}
	}
}