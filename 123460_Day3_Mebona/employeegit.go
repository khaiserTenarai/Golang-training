package main

import "fmt"

type Staff struct {
	ID     int
	Name   string
	Role   string
	Salary int
}

var team []Staff

func main() {
	for {
		fmt.Println("\n===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Search Employee")
		fmt.Println("6. Exit")

		var opt int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&opt)

		switch opt {

		case 1:
			var s Staff

			fmt.Print("Enter ID: ")
			fmt.Scan(&s.ID)

			fmt.Print("Enter Name: ")
			fmt.Scan(&s.Name)

			fmt.Print("Enter Role: ")
			fmt.Scan(&s.Role)

			fmt.Print("Enter Salary: ")
			fmt.Scan(&s.Salary)

			team = append(team, s)

			fmt.Println("Employee added successfully.")

		case 2:
			if len(team) == 0 {
				fmt.Println("No employees found.")
				continue
			}

			fmt.Println("\n----- Employee Directory -----")

			for _, s := range team {
				fmt.Printf(
					"ID: %d | Name: %s | Role: %s | Salary: %d\n",
					s.ID, s.Name, s.Role, s.Salary,
				)
			}

		case 3:
			var targetID int
			fmt.Print("Enter ID to update: ")
			fmt.Scan(&targetID)

			found := false

			for i := range team {
				if team[i].ID == targetID {

					fmt.Print("Enter New Name: ")
					fmt.Scan(&team[i].Name)

					fmt.Print("Enter New Role: ")
					fmt.Scan(&team[i].Role)

					fmt.Print("Enter New Salary: ")
					fmt.Scan(&team[i].Salary)

					fmt.Println("Employee updated successfully.")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 4:
			var targetID int
			fmt.Print("Enter ID to delete: ")
			fmt.Scan(&targetID)

			found := false

			for i, s := range team {
				if s.ID == targetID {

					team = append(team[:i], team[i+1:]...)

					fmt.Println("Employee deleted successfully.")
					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 5:
			var targetID int
			fmt.Print("Enter Employee ID to search: ")
			fmt.Scan(&targetID)

			found := false

			for _, s := range team {
				if s.ID == targetID {
					fmt.Println("\nEmployee Found:")
					fmt.Printf("ID: %d\n", s.ID)
					fmt.Printf("Name: %s\n", s.Name)
					fmt.Printf("Role: %s\n", s.Role)
					fmt.Printf("Salary: %d\n", s.Salary)

					found = true
					break
				}
			}

			if !found {
				fmt.Println("Employee not found.")
			}

		case 6:
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}