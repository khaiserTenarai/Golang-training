package dao

import "fmt"

var employeeId int
var employeeName string
var employeeSalary int

func AddEmployee(id int, name string, salary int) {
	employeeId = id
	employeeName = name
	employeeSalary = salary
}
func SearchEmployee(id int) bool {
	if employeeId == id {
		return true
	}
	return false
}
func DisplayEmployees() {
	fmt.Println("Employee id: ", employeeId)
	fmt.Println("Employee Name: ", employeeName)
	fmt.Println("Employee salary: ", employeeSalary)
}
func DeleteEmployee(id int) bool {
	if employeeId == id {
		employeeId = 0
		employeeName = ""
		employeeSalary = 0
		return true
	}
	return false
}
