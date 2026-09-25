package main

import "fmt"

func main() {

	fmt.Println("\n*****************************************")
	fmt.Println("11. Implement employee lookup using a map. ")
	fmt.Println("******************************************")

	employees := map[int]string{
		101: "Vimal",
		102: "Aditya",
		103: "Adi",
	}

	var choice int

	for {
		fmt.Println("\n--- EMPLOYEE MAP LOOKUP ---")
		fmt.Println("1. Search Employee by ID")
		fmt.Println("2. Add New Employee")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var searchID int
			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&searchID)

			name, exists := employees[searchID]

			if exists {
				fmt.Println("Result -> Employee Name:", name)
			} else {
				fmt.Println("Result -> Employee not found for ID:", searchID)
			}

		case 2:
			var newID int
			var newName string

			fmt.Print("Enter New Employee ID: ")
			fmt.Scan(&newID)

			name, exists := employees[newID]
			if exists {
				fmt.Println("Error: ID already belongs to", name)
			} else {
				fmt.Print("Enter Employee Name: ")
				fmt.Scan(&newName)

				employees[newID] = newName
				fmt.Println("Employee added successfully!")
			}

		case 3:
			fmt.Println("Exiting application. Goodbye!")
			return

		default:
			fmt.Println("Invalid option! Please try again.")
		}
	}
}