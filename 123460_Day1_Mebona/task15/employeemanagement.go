package main

import (
	"fmt"
)

func main() {
	// Using parallel slices 
	var ids []string
	var names []string

	for {
		fmt.Println("\nEmployee Management")
		fmt.Println("1. Add employee")
		fmt.Println("2. Search employee")
		fmt.Println("3. Display employees")
		fmt.Println("4. Delete employee")
		fmt.Println("5. Exit")
		fmt.Print("Select an option: ")

		var choice string
		fmt.Scan(&choice)

		switch choice {
		case "1":
			var id, name string
			fmt.Print("Enter Employee ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Employee Name (single word): ")
			fmt.Scan(&name)

		
			ids = append(ids, id)
			names = append(names, name)
			fmt.Println("Employee added successfully!")

		case "2":
			var searchID string
			fmt.Print("Enter Employee ID to search: ")
			fmt.Scan(&searchID)

			found := false
			for i := 0; i < len(ids); i++ {
				if ids[i] == searchID {
					fmt.Printf("Found -> ID: %s, Name: %s\n", ids[i], names[i])
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case "3":
			fmt.Println("\n--- All Employees ---")
			if len(ids) == 0 {
				fmt.Println("No employees in the system.")
			} else {
				for i := 0; i < len(ids); i++ {
					fmt.Printf("ID: %s | Name: %s\n", ids[i], names[i])
				}
			}

		case "4":
			var deleteID string
			fmt.Print("Enter Employee ID to delete: ")
			fmt.Scan(&deleteID)

			found := false
			for i := 0; i < len(ids); i++ {
				if ids[i] == deleteID {
					
					ids = append(ids[:i], ids[i+1:]...)
					names = append(names[:i], names[i+1:]...)
					
					fmt.Println("Employee deleted successfully!")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case "5":
			fmt.Println("Exiting application...")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}