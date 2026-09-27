package main

import "fmt"

type Address struct{
	City string
	State string
	Pincode string
}

type Employee struct{
	ID int
	Name string
	Address Address
}

func main(){
	emp := Employee{
		ID: 1,
		Name: "Sanskriti",
		Address: Address{
			City: "Bangalore",
			State: "Karnataka",
			Pincode: "560076",
		},
	}

	fmt.Println("ID",emp.ID)
	fmt.Println("Name",emp.Name)
}