package main

import "fmt"

type Employee struct{
	ID int 
	Name string
	Salary float64
	Department string
}

func (e Employee) Display(){
	fmt.Println("ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Salary:", e.Salary)
	fmt.Println("Department:", e.Department)
}