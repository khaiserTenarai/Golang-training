package controller

import (
	"fmt"
	"task15/task15/service"
)

func AddEmployee() {
	var id int
	var name string
	var salary int

	fmt.Println("Enter Employee ID: ")
	fmt.Scan(&id)

	fmt.Println("Enter Employee Name: ")
	fmt.Scan(&name)

	fmt.Println("Enter Employee Salary: ")
	fmt.Scan(&salary)

	service.AddEmployee(id, name, salary)

	fmt.Println("Employee added successfully")
}

func SearchEmployee() {
	var id int
	fmt.Println("Enter Employee ID to search: ")
	fmt.Scan(&id)

	if service.SearchEmployee(id) {
		fmt.Println("Employee Found")
	} else {
		fmt.Println("Employee not found")
	}
}

func DisplayEmployees() {
	service.DisplayEmployees()
}

func DeleteEmployee() {
	var id int
	fmt.Println("Enter Employee ID to delete: ")
	fmt.Scan(&id)

	if service.DeleteEmployee(id) {
		fmt.Println("Employee deleted successfully")
	} else {
		fmt.Println("Employee not found")
	}
}
