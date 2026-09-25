package main

import "fmt"

var ids []int
var names []string
var salaries []int
var employeeMap = make(map[int]int)

func addEmployee() {
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
	employeeMap[id] = len(ids) - 1

	fmt.Println("Employee added successfully")
}
func searchEmployee() {
	var id int
	fmt.Println("Enter the ID: ")
	fmt.Scan(&id)
	index, found := employeeMap[id]
	if found {
		fmt.Println("ID :", ids[index])
		fmt.Println("Name :", names[index])
		fmt.Println("Salary :", salaries[index])
	} else {
		fmt.Println("Employee not found")
	}
}
func displayEmployees() {
	if len(ids) == 0 {
		fmt.Println("Employee not found")
	}
	for i := 0; i < len(ids); i++ {
		fmt.Println("ID :", ids[i])
		fmt.Println("Name :", names[i])
		fmt.Println("Salary :", salaries[i])
	}
}
func deleteEmployee() {
	var id int
	fmt.Println("Enter ID to delete: ")
	fmt.Scan(&id)

	index, found := employeeMap[id]
	if found {
		ids = append(ids[:index], ids[index+1:]...)
		names = append(names[:index], names[index+1:]...)
		salaries = append(salaries[:index], salaries[index+1:]...)
		delete(employeeMap, id)
		for i := index; i < len(ids); i++ {
			employeeMap[ids[i]] = i
		}
		fmt.Println("Employee deleted successfully")
	} else {
		fmt.Println("Employee not found")
	}
}
func main() {
	var choice int
	for {
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Update Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")

		fmt.Println("Enter your choice: ")
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
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
