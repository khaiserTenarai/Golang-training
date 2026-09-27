package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Salary    float64
}

func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e Employee) TryToUpdateName(first string, last string) {
	e.FirstName = first
	e.LastName = last
}

func main() {
	emp := Employee{
		FirstName: "Jane",
		LastName:  "Doe",
		Salary:    100000.0,
	}

	fmt.Println("Full Name:", emp.FullName())

	emp.TryToUpdateName("John", "Smith")
	fmt.Println("After Update Try:", emp.FirstName) 
}