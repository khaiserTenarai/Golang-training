package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Age  int
}

// Value receiver
func (e Employee) Display() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
}

func main() {
	employee := Employee{
		ID:   101,
		Name: "Rajesh",
		Age:  25,
	}

	employee.Display()
}
