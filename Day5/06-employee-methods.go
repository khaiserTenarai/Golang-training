package main

import "fmt"



type Employee struct {
 ID int
 Name string
 Salary float64
}

// Method
func (e Employee) Display() {
 fmt.Println("ID:", e.ID)
 fmt.Println("Name:", e.Name)
 fmt.Println("Salary:", e.Salary)
}

func main() {

 emp := Employee{
  ID: 101,
  Name: "Ganesh",
  Salary: 40000,
 }

 emp.Display()
}

/*
PS C:\Training\Go Lang\Day_5\124772_Day5(Go)_Assignment_Reddem_Ganesh_Reddy> go run .\06-employee-methods.go
ID: 101
Name: Ganesh
Salary: 40000
*/