package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Age  int
}

type EmployeeRepository interface {
	Add(employee Employee)
	GetByID(id int) Employee
	GetAll() []Employee
}

func main() {
	var repo EmployeeRepository
	_ = repo

	fmt.Println("EmployeeRepository interface created")
}
