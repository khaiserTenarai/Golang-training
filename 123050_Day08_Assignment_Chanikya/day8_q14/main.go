package main

import (
	"fmt"
)

var employees []string

func addEmployee(name string) {
	employees = append(employees, name)
}

func processEmployee(name string) {
	fmt.Println("Employee:", name)
	fmt.Println("Employee:", name)
}

func main() {
	addEmployee("max")
	addEmployee("Rahul")

	for i := 0; i < len(employees); i++ {
		processEmployee(employees[i])
	}

	fmt.Println("Salary:", 50000)
}
