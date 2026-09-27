package main

import "fmt"

type Employee struct {
 Name string
 Salary float64
}

// Value receiver
func (e Employee) IncreaseSalary() {
 e.Salary = e.Salary + 5000
 fmt.Println("Inside method:", e.Salary)
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
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\08-value-reciever.go
Before: 40000
Inside method: 45000
After: 40000
*/