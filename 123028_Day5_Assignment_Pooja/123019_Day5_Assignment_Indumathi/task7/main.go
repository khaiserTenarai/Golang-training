package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Salary    float64
	IsActive  bool
}

func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * (percent / 100.0)
}

func (e *Employee) Deactivate() {
	e.IsActive = false
}

func main() {
	emp := Employee{
		FirstName: "Jane",
		LastName:  "Doe",
		Salary:    100000.0,
		IsActive:  true,
	}

	fmt.Println("Before:", emp.Salary, emp.IsActive) 

	emp.GiveRaise(10)
	emp.Deactivate()

	fmt.Println("After: ", emp.Salary, emp.IsActive)
}