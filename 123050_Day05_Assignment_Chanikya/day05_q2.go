package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Department struct {
	Name string
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
		Name: "Ram",
		Address: Address{
			City:  "Bangalore",
			State: "Karnataka",
		},
		Department: Department{
			Name: "IT",
		},
	}

	fmt.Println(employee.Name)
	fmt.Println(employee.Address.City)
	fmt.Println(employee.Department.Name)
}
