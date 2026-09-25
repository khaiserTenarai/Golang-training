package main

import (
	"fmt"

	"employeemgmt/employee"
)

var employees []employee.Employee

func main() {
	for {
		fmt.Println("\n1. Add 2. Search 3. Display 4. Delete 5. Exit")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var e employee.Employee
			fmt.Print("Enter ID, Name, Dept, Salary: ")
			fmt.Scan(&e.ID, &e.Name, &e.Dept, &e.Salary)
			employees = append(employees, e)
			fmt.Println("Employee added!")

		case 2:
			var id string
			fmt.Print("Enter ID to search: ")
			fmt.Scan(&id)

			found := false
			for _, e := range employees {
				if e.ID == id {
					fmt.Printf("Found -> ID: %s, Name: %s, Dept: %s, Salary: %.2f\n", e.ID, e.Name, e.Dept, e.Salary)
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Employee not found.")
			}

		case 3:
			if len(employees) == 0 {
				fmt.Println("No employees to show.")
				continue
			}
			for _, e := range employees {
				fmt.Printf("ID: %s | Name: %s | Dept: %s | Salary: %.2f\n", e.ID, e.Name, e.Dept, e.Salary)
			}

		case 4:
			var id string
			fmt.Print("Enter ID to delete: ")
			fmt.Scan(&id)

			var updatedList []employee.Employee
			for _, e := range employees {
				if e.ID != id {
					updatedList = append(updatedList, e)
				}
			}
			employees = updatedList
			fmt.Println("Delete complete.")

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}