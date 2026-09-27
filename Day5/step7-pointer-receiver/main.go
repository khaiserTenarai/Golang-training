package main

import "fmt"

type Employee struct {
	Name   string
	Salary float64
}


func (e *Employee) GiveRaise(percent float64) {
	e.Salary += e.Salary * percent / 100
}

type EmployeeStatus struct {
	Employee
	IsActive bool
}

func (s *EmployeeStatus) Deactivate() {
	s.IsActive = false
}

func main() {
	e := Employee{Name: "Ray", Salary: 60000}

	fmt.Println("Before raise:", e.Salary)
	e.GiveRaise(10) // note:Go automatically takes &e 
	fmt.Println("After raise: ", e.Salary)

	status := EmployeeStatus{Employee: e, IsActive: true}
	fmt.Println("\nBefore deactivate:", status.IsActive)
	status.Deactivate()
	fmt.Println("After deactivate: ", status.IsActive)

	fmt.Println("\nWhy pointer receiver: without it, GiveRaise would modify")
	fmt.Println("a COPY of the Employee, and the original e.Salary would never change.")
}
