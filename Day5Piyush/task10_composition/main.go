package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Introduce() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}

type Employee struct {
	Person
	EmployeeID int
	Salary     float64
}

func main() {

	employee := Employee{
		Person: Person{
			Name: "Piyush",
			Age:  24,
		},
		EmployeeID: 101,
		Salary:     50000,
	}

	employee.Introduce()

	fmt.Println("Employee ID:", employee.EmployeeID)
	fmt.Println("Salary:", employee.Salary)
}