package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	ID         int
	Department string
}

func main() {
	employee := Employee{
		Person: Person{
			Name: "Ravi",
			Age:  25,
		},
		ID:         101,
		Department: "IT",
	}

	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.ID)
	fmt.Println(employee.Department)
}
