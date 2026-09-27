package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Display() {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age :", p.Age)
}

// Composition: Employee has Person instead of inheriting from it
type Employee9 struct {
	Person Person  
	ID      int
	Salary  float64
}

func main() {
 
	fmt.Println("\n***********************************************")
	fmt.Println("10. Implement composition instead of inheritance")
	fmt.Println("************************************************")

	emp := Employee9{
		Person: Person{
			Name: "Vimal",
			Age:  23,
		},
		ID:     101,
		Salary: 500000,
	}

	fmt.Println("ID  :", emp.ID)
	emp.Person.Display()

}