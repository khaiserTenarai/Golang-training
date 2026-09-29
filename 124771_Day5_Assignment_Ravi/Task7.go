package main

import "fmt"

type Employee struct {
	ID    int
	Name  string
	Age   int
	City  string
	State string
}

func (e Employee) Display() {
	fmt.Println("ID:", e.ID)
	fmt.Println("Name:", e.Name)
	fmt.Println("Age:", e.Age)
	fmt.Println("City:", e.City)
}

func main() {
	employee := Employee{
		ID:    101,
		Name:  "Rajesh",
		Age:   25,
		City:  "Bangalore",
		State: "Karnataka",
	}

	employee.Display()
}
