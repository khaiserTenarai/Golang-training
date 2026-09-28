package main

import "fmt"

type Employee6 struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Salary float64
}

func (employee Employee6) display() {
	fmt.Println("ID: ", employee.ID)
	fmt.Println("Name: ", employee.Name)
	fmt.Println("Email: ", employee.Email)
	fmt.Println("Age: ", employee.Age)
	fmt.Println("Salary: ", employee.Salary)
}
func main() {
	employee := Employee6{
		ID:     1,
		Name:   "Abirami",
		Email:  "abi@gmail.com",
		Age:    22,
		Salary: 80000,
	}
	employee.display()
}
