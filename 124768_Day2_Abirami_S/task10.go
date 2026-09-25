package main

import "fmt"

func main() {
	ids := []int{}
	names := []string{}
	salaries := []int{}
	for {
		var choice int
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id int
			var name string
			var salary int

			fmt.Println("Enter ID: ")
			fmt.Scan(&id)

			fmt.Println("Enter Name: ")
			fmt.Scan(&name)

			fmt.Println("Enter Salary: ")
			fmt.Scan(&salary)

			ids = append(ids, id)
			names = append(names, name)
			salaries = append(salaries, salary)

			fmt.Println("Employee added successfully")
		case 2:
			for i := 0; i < len(ids); i++ {
				fmt.Println("ID :", ids[i])
				fmt.Println("Name :", names[i])
				fmt.Println("Salary :", salaries[i])
			}
		case 3:
			var id int
			fmt.Println("Enter the ID: ")
			fmt.Scan(&id)
			for i := 0; i < len(ids); i++ {
				if ids[i] == id {
					fmt.Println("Enter New Name: ")
					fmt.Scan(&names[i])
					fmt.Println("Enter New Salary: ")
					fmt.Scan(&salaries[i])
					fmt.Println("Employee details updated successfully")
				}
			}
		case 4:
			var id int
			fmt.Println("Enter ID to delete: ")
			fmt.Scan(&id)

			for i := 0; i < len(ids); i++ {
				if ids[i] == id {
					ids = append(ids[:i], ids[i+1:]...)
					names = append(names[:i], names[i+1:]...)
					salaries = append(salaries[:i], salaries[i+1:]...)
					fmt.Println("Employee deleted successfully")
					break
				}
			}
		case 5:
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
