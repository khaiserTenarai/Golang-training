package main

import "fmt"

type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Address Address // Composition
}

func main() {
	employee := Employee{
		Name: "Ram",
		Address: Address{
			City:  "Bangalore",
			State: "Karnataka",
		},
	}

	fmt.Println(employee.Name)
	fmt.Println(employee.Address.City)
	fmt.Println(employee.Address.State)
}
