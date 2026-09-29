package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

// EmployeeRepository interface
type EmployeeRepository interface {
	Add(employee Employee)
	GetByID(id int) Employee
	GetAll() []Employee
	Delete(id int)
}

func main() {
	fmt.Println("EmployeeRepository interface created")
}
