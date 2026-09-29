package main

import "fmt"

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Department struct {
	Name string `json:"name"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
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

	fmt.Println(employee)
}
