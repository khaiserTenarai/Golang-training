package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Pincode string `json:"pincode"`
}

type Department struct {
	Name string `json:"name"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Age        int        `json:"age"`
	Salary     float64    `json:"salary"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
}

func main() {
	e := Employee{
		ID:      1,
		Name:    "Ray",
		Email:   "ray@example.com",
		Age:     26,
		Salary:  60000,
		Address: Address{City: "Bengaluru", Pincode: "560001"},
		Department: Department{
			Name: "Engineering",
		},
	}


	data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}

	fmt.Println("Employee marshaled to JSON:")
	fmt.Println(string(data))
}
