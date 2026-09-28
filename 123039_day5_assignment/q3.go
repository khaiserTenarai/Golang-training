package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
}

type Department struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Age        int        `json:"age"`
	Salary     float64    `json:"salary"`
	Phone      string     `json:"phone"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Swathi",
		Email:  "swathi@gmail.com",
		Age:    25,
		Salary: 50000,
		Phone:  "9876543210",

		Address: Address{
			Street:  "MG Road",
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: 560001,
		},

		Department: Department{
			Name:    "IT",
			Manager: "Rahul",
		},
	}

	jsonData, err := json.MarshalIndent(employee, "", "    ")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}