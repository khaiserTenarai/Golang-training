package main

import "fmt"

var id int
var name string
var salary float64

func createEmployee() {
	fmt.Println("Enter employee id: ")
	fmt.Scan(&id)
	fmt.Println("Enter employee name: ")
	fmt.Scan(&name)
	fmt.Println("Enter employee salary: ")
	fmt.Scan(&salary)
	fmt.Println("Employee created")
}
func viewEmployee() {
	fmt.Println("Id: ", id)
	fmt.Println("Name: ", name)
	fmt.Println("Salary: ", salary)
}
func updateEmployee() {
	fmt.Println("Enter employee id: ")
	fmt.Scan(&id)
	fmt.Println("Enter employee's new name: ")
	fmt.Scan(&name)
	fmt.Println("Enter employee's new salary: ")
	fmt.Scan(&salary)
	fmt.Println("Employee updated")
}
func deleteEmployee() {
	id = 0
	name = ""
	salary = 0
	fmt.Println("Employee deleted")
}
func main() {
	createEmployee()
	viewEmployee()
	updateEmployee()
	deleteEmployee()
}
