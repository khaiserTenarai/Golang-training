package main

import (
	"fmt"
)

type Staff struct {
	ID   int
	Name string
	Role string
}

var team []Staff

func main() {
	for {
		fmt.Println("\n1. Add | 2. View | 3. Update | 4. Delete | 5. Exit")
		var opt int
		fmt.Scan(&opt)

		switch opt {
		case 1:
			var s Staff
			fmt.Print("Enter ID, Name, and Role: ")
			fmt.Scan(&s.ID, &s.Name, &s.Role)
			
			team = append(team, s)
			fmt.Println("Added successfully.")

		case 2:
			fmt.Println("Team Directory:")
			for _, s := range team {
				fmt.Printf("ID: %d | Name: %s | Role: %s\n", s.ID, s.Name, s.Role)
			}

		case 3:
			var targetID int
			fmt.Print("Enter ID to update: ")
			fmt.Scan(&targetID)

			for i := range team {
				if team[i].ID == targetID {
					fmt.Print("Enter New Name and Role: ")
					fmt.Scan(&team[i].Name, &team[i].Role)
					fmt.Println("Updated successfully.")
					break
				}
			}

		case 4:
			var targetID int
			fmt.Print("Enter ID to delete: ")
			fmt.Scan(&targetID)

			for i, s := range team {
				if s.ID == targetID {
					
					team = append(team[:i], team[i+1:]...)
					fmt.Println("Deleted successfully.")
					break
				}
			}

		case 5:
			return
			
		default:
			fmt.Println("Invalid option.")
		}
	}
}