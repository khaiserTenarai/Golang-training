package main

import "fmt"


type Address struct {
	City    string
	Pincode string
}

type Department struct {
	Name string
}

type Employee struct {
	ID         int
	Name       string
	Email      string
	Age        int
	Salary     float64
	Address    Address    // taking nested struct
	Department Department //same hereee
}

func main() {
	e := Employee{
		ID:     1,
		Name:   "Ray",
		Email:  "ray@example.com",
		Age:    26,
		Salary: 60000,
		Address: Address{
			City:    "Bengaluru",
			Pincode: "560001",
		},
		Department: Department{
			Name: "Engineering",
		},
	}

	fmt.Println("Employee with nested Address & Department:")
	fmt.Printf("%+v\n", e)

	// Accessing nested fields
	fmt.Println("City:", e.Address.City)
	fmt.Println("Department:", e.Department.Name)
}
