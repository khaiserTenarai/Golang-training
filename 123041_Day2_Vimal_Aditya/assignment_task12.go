package main

import "fmt"

type Profile struct {
	ID   int
	Name string
}

func main() {

	fmt.Println("\n********************************************")
	fmt.Println("12. Create a menu-driven program using switch.")
	fmt.Println("**********************************************")

	var profiles []Profile
	var choice int
	var continueChoice string

	for {
		fmt.Println("\n************** MAIN MENU **************")
		fmt.Println("1. Create Profile")
		fmt.Println("2. View Profile by ID")
		fmt.Println("3. Delete Profile")
		fmt.Println("4. Update Profile")
		fmt.Println("5. View All Profiles")
		fmt.Println("6. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			var id int
			var name string

			fmt.Print("Enter ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			profiles = append(profiles, Profile{ID: id, Name: name})
			fmt.Println("Profile created successfully!")

		case 2:
			if len(profiles) == 0 {
				fmt.Println("No profiles available.")
			} else {
				var searchID int
				fmt.Print("Enter ID to view: ")
				fmt.Scan(&searchID)

				found := false
				for i := 0; i < len(profiles); i++ {
					if profiles[i].ID == searchID {
						fmt.Println("Found Profile -> ID:", profiles[i].ID, "| Name:", profiles[i].Name)
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Profile not found.")
				}
			}

		case 3:
			if len(profiles) == 0 {
				fmt.Println("No profiles to delete.")
			} else {
				var deleteID int
				fmt.Print("Enter ID to delete: ")
				fmt.Scan(&deleteID)

				found := false
				for i := 0; i < len(profiles); i++ {
					if profiles[i].ID == deleteID {
						profiles = append(profiles[:i], profiles[i+1:]...)
						fmt.Println("Profile deleted successfully!")
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Profile not found.")
				}
			}

		case 4:
			if len(profiles) == 0 {
				fmt.Println("No profiles to update.")
			} else {
				var updateID int
				fmt.Print("Enter ID to update: ")
				fmt.Scan(&updateID)

				found := false
				for i := 0; i < len(profiles); i++ {
					if profiles[i].ID == updateID {
						var newName string
						fmt.Print("Enter new name: ")
						fmt.Scan(&newName)

						profiles[i].Name = newName
						fmt.Println("Profile updated successfully!")
						found = true
						break
					}
				}
				if !found {
					fmt.Println("Profile not found.")
				}
			}

		case 5:
			if len(profiles) == 0 {
				fmt.Println("No profiles found.")
			} else {
				fmt.Println("\n--- All Profiles ---")
				for i := 0; i < len(profiles); i++ {
					fmt.Println("ID:", profiles[i].ID, "| Name:", profiles[i].Name)
				}
			}

		case 6:
			fmt.Println("Thank you!")
			return

		default:
			fmt.Println("Wrong choice!")
		}

		fmt.Print("\nDo you want to continue? (yes/no): ")
		fmt.Scan(&continueChoice)

		if continueChoice != "yes" {
			fmt.Println("Thank you!")
			break
		}
	}
}