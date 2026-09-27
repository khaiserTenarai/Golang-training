package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Department struct {
	Name    string
	Manager string
	Floor   int
}

type Employee struct {
	ID       int
	FullName string
	Location Address
	Role     Department
}

func (e Employee) PrintSummary() {
	fmt.Printf("Employee: %s (ID: %d)\n", e.FullName, e.ID)
	fmt.Printf("Location: %s, %s\n", e.Location.City, e.Location.State)
	fmt.Printf("Department: %s\n", e.Role.Name)
}

func (e *Employee) Relocate(newCity string, newState string) {
	e.Location.City = newCity
	e.Location.State = newState
}

func (e *Employee) ChangeDepartment(newDepartment string, newManager string) {
	e.Role.Name = newDepartment
	e.Role.Manager = newManager
}

func main() {
	emp := Employee{
		ID:       301,
		FullName: "Lakshmi Shibu",
		Location: Address{
			Street:  "123 MG Road",
			City:    "Munnar",
			State:   "Kerala",
			ZipCode: "59715",
		},
		Role: Department{
			Name:    "Logistics",
			Manager: "Rahul",
			Floor:   1,
		},
	}

	emp.PrintSummary()

	fmt.Println("\nProcessing transfer...")

	emp.Relocate("Bangalore", "BLR")
	emp.ChangeDepartment("Operations", "Laiya")

	fmt.Println("\nUpdated Profile:")
	emp.PrintSummary()
}