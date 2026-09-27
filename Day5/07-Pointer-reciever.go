package main

import "fmt"

type Employee struct {
 Name string
 Salary float64
}

func (e *Employee) IncreaseSalary() {
 e.Salary = e.Salary + 5000
}

func main() {

 emp := Employee{
  Name: "Ganesh",
  Salary: 40000,
 }

 fmt.Println("Before:", emp.Salary)

 emp.IncreaseSalary()

 fmt.Println("After:", emp.Salary)
}

/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\07-Pointer-reciever.go
Before: 40000
After: 45000
*/