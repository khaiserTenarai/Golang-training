package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	Pincode string
}

type Department struct {
	Name    string
	Team    string
	Manager string
}

type Employee struct {
	ID         int
	Name       string
	Address    Address
	Department Department
}

func main() {
	employee := Employee{
		ID:   101,
		Name: "John",

		Address: Address{
			Street:  "MG Road",
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: "560001",
		},

		Department: Department{
			Name:    "IT",
			Team:    "Development",
			Manager: "David",
		},
	}

	fmt.Println("Employee:", employee.Name)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("Department:", employee.Department.Name)
}