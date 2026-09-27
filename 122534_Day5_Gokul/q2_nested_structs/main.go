/*
Day 5, Q2. Create nested Address and Department structs.
*/
package main

import "fmt"

type Address struct {
	City  string
	State string
	Pin   string
}

type Department struct {
	Name string
	Code string
}

type Employee struct {
	Name       string
	Address    Address    
	Department Department 
}

func main() {
	gokul := Employee{
		Name: "Gokul",
		Address: Address{
			City:  "Bengaluru",
			State: "Karnataka",
			Pin:   "560001",
		},
		Department: Department{
			Name: "Data Engineering",
			Code: "DE-04",
		},
	}

	fmt.Println("Name      :", gokul.Name)
	fmt.Println("City      :", gokul.Address.City)
	fmt.Println("State     :", gokul.Address.State)
	fmt.Println("Department:", gokul.Department.Name)
}
