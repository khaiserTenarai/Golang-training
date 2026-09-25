package main

import "fmt"

type Employee struct {
	ID     string
	Name   string
	Dept   string
	Salary float64
}

func main() {
	var employees []Employee
	var activityLog []string

	for {
		fmt.Println("\n1. Add | 2. Search | 3. Display | 4. Delete | 5. Activity Log | 6. Exit")
		fmt.Print("Choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var e Employee
			fmt.Print("Enter ID, Name, Dept, Salary: ")
			fmt.Scan(&e.ID, &e.Name, &e.Dept, &e.Salary)
			exists := false
			for _, emp := range employees {
				if emp.ID == e.ID {
					exists = true
					break
				}
			}
			if exists {
				fmt.Println("Employee ID already exists.")
				continue
			}

			employees = append(employees, e)
			activityLog = append(activityLog, "Added employee "+e.ID)
			fmt.Println("Employee added!")

		case 2:
			var id string
			fmt.Print("Enter ID to search: ")
			fmt.Scan(&id)

			found := false
			for _, e := range employees {
				if e.ID == id {
					fmt.Printf("ID: %s, Name: %s, Dept: %s, Salary: %.2f\n", e.ID, e.Name, e.Dept, e.Salary)
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

			var updated []Employee
			deleted := false
			for _, e := range employees {
				if e.ID != id {
					updated = append(updated, e)
				} else {
					deleted = true
				}
			}

			if deleted {
				employees = updated
				activityLog = append(activityLog, "Deleted employee "+id)
				fmt.Println("Employee deleted!")
			} else {
				fmt.Println("Employee not found.")
			}

		case 5:
			if len(activityLog) == 0 {
				fmt.Println("No activity recorded yet.")
				continue
			}
			for _, log := range activityLog {
				fmt.Println("-", log)
			}

		case 6:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}