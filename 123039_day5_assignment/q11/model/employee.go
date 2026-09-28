package model

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Email  string
	Age    int
	Salary float64
	Phone  string
}

func (e Employee) Display() {
	fmt.Println("----------------------------")
	fmt.Println("Employee ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Email:", e.Email)
	fmt.Println("Age:", e.Age)
	fmt.Println("Salary:", e.Salary)
	fmt.Println("Phone:", e.Phone)
	fmt.Println("----------------------------")
}

func (e *Employee) IncreaseSalary(amount float64) {
	e.Salary = e.Salary + amount
}