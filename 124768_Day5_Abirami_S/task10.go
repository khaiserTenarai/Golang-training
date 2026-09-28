package main

import "fmt"

type Address1 struct {
	City    string
	State   string
	Pincode string
}
type Employee10 struct {
	ID      int
	Name    string
	Address Address1
}

func main() {
	employee := Employee10{
		ID:   101,
		Name: "Abirami",

		Address: Address1{
			City:    "Banglore",
			State:   "Karnataka",
			Pincode: "560100",
		},
	}
	fmt.Println("ID: ", employee.ID)
	fmt.Println("Name: ", employee.Name)
	fmt.Println("City: ", employee.Address.City)
	fmt.Println("State: ", employee.Address.State)
	fmt.Println("Pincode: ", employee.Address.Pincode)
}
