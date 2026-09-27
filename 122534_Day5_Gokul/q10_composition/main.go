package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

type Employee struct {
	Person 
	Salary float64
	Dept   string
}

func main() {
	gokul := Employee{
		Person: Person{Name: "Gokul", Age: 27},
		Salary: 65000,
		Dept:   "Data Engineering",
	}


	fmt.Println(gokul.Greet())
	fmt.Println("Name:", gokul.Name, "Age:", gokul.Age)
	fmt.Println("Dept:", gokul.Dept, "Salary:", gokul.Salary)
}
