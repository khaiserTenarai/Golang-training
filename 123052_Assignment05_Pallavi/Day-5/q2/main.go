package main

import "fmt"

// Address represents employee address
type Address struct {
	Street string
	City   string
	State  string
	Pincode int
}

// Department represents employee department
type Department struct {
	ID   int
	Name string
}

// Employee contains nested structs
type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Address    Address
	Department Department
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Pallavi",
		Email:  "pallavi@gmail.com",
		Age:    22,
		Salary: 29000,

		Address: Address{
			Street: "MG Road",
			City:   "Bangalore",
			State:  "Karnataka",
			Pincode: 560001,
		},

		Department: Department{
			ID:   10,
			Name: "IT",
		},
	}

	fmt.Println("Employee Name:", employee.Name)

	fmt.Println("\nAddress:")
	fmt.Println("Street:", employee.Address.Street)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Pincode:", employee.Address.Pincode)

	fmt.Println("\nDepartment:")
	fmt.Println("ID:", employee.Department.ID)
	fmt.Println("Name:", employee.Department.Name)
}