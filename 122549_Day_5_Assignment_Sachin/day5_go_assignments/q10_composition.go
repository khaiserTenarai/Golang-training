// 10. Implement composition instead of inheritance.

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Speak() {
	fmt.Printf("Hi, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

type Employee struct {
	Person
	EmployeeID int
	Position   string
}

func main() {
	emp := Employee{
		Person:     Person{Name: "sachin", Age: 28},
		EmployeeID: 505,
		Position:   "Software Engineer",
	}

	emp.Speak()
	fmt.Println("Employee ID:", emp.EmployeeID)
	fmt.Println("Position:", emp.Position)
}
