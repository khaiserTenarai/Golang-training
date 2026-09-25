package main

import "fmt"

var ids [100]int
var names [100]string
var roles [100]string

var count int

func addEmployee() {

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&ids[count])

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&names[count])

	fmt.Print("Enter Employee Role: ")
	fmt.Scan(&roles[count])

	count++

	fmt.Println("Employee added successfully!")
}

func searchEmployee() {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	for i := 0; i < count; i++ {

		if ids[i] == id {
			fmt.Println("Employee ID:", ids[i])
			fmt.Println("Name:", names[i])
			fmt.Println("Role:", roles[i])
			return
		}
	}

	fmt.Println("Employee not found")
}

func displayEmployees() {

	if count == 0 {
		fmt.Println("No employees found")
		return
	}

	for i := 0; i < count; i++ {
		fmt.Println("----------------")
		fmt.Println("Employee ID:", ids[i])
		fmt.Println("Name:", names[i])
		fmt.Println("Role:", roles[i])
	}
}

func deleteEmployee() {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	for i := 0; i < count; i++ {

		if ids[i] == id {

			for j := i; j < count-1; j++ {
				ids[j] = ids[j+1]
				names[j] = names[j+1]
				roles[j] = roles[j+1]
			}

			count--

			fmt.Println("Employee deleted successfully!")
			return
		}
	}

	fmt.Println("Employee not found")
}

func main() {

	for {

		fmt.Println("\nEmployee Management")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee")
		fmt.Println("3. Display Employees")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		var choice int

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			addEmployee()

		case 2:
			searchEmployee()

		case 3:
			displayEmployees()

		case 4:
			deleteEmployee()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
