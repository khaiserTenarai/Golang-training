package main

import "fmt"


type Employee6 struct {
	ID         int
	Name       string
	Salary     float64
}

func (e Employee6) DisplayDetails() {
	fmt.Println("ID        :", e.ID)
	fmt.Println("Name      :", e.Name)
	fmt.Println("Salary    :", e.Salary)
}

func (e *Employee6) UpdateSalary(newSalary float64) {
	e.Salary = newSalary
}

func main() {

	fmt.Println("\n*****************************")
	fmt.Println("6. Create Employee methods. ")
	fmt.Println("****************************")

	emp := Employee6{
		ID:     101,
		Name:   "Vimal Aditya",
		Salary: 500000,
	}

	emp.DisplayDetails()

	emp.UpdateSalary(600000)

	fmt.Println("\nAfter Salary Update:")
	fmt.Println("New Salary:", emp.Salary)
}