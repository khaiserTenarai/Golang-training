package main

import "fmt"

type Address struct {
	City    string
	State   string
	Pincode string
}

type Employee struct {
	ID      int
	Name    string
	Address Address
}

func main() {
	employee := Employee{
		ID:   101,
		Name: "fammy",
		Address: Address{
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: "18631",
		},
	}

	fmt.Println(employee)
}
