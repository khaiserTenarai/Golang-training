package main

import "fmt"

// Person is the shared, reusable piece.
type Person struct {
	Name  string
	Email string
	Age   int
}

func (p Person) Greet() string {
	return "Hi, I'm " + p.Name
}

type Employee struct {
	Person // embedded, not inherited
	ID     int
	Salary float64
}

func main() {
	e := Employee{
		Person: Person{Name: "Ray", Email: "ray@example.com", Age: 26},
		ID:     1,
		Salary: 60000,
	}

	fmt.Println("Name :", e.Name)
	fmt.Println("Email :", e.Email)

	fmt.Println("Greet :", e.Greet())

	fmt.Println("Explicit access:", e.Person.Name)

	fmt.Println("\nWhy this matters: Employee doesn't inherit from Person")
	fmt.Println("in the OOP sense. It simply CONTAINS a Person value, and")
	fmt.Println("Go's promotion rules make that contained value's fields")
	fmt.Println("and methods reachable as if they belonged to Employee.")
}
