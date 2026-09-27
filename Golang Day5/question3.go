package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street_address"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
}

type Department struct {
	Name    string `json:"department_name"`
	Manager string `json:"manager_name"`
	Floor   int    `json:"floor_number"`
}

type Employee struct {
	ID       int        `json:"employee_id"`
	FullName string     `json:"full_name"`
	Location Address    `json:"location"`
	Role     Department `json:"role"`
}

func main() {
	emp := Employee{
		ID:       101,
		FullName: "Lakshmi Shibu",
		Location: Address{
			Street:  "MG Road",
			City:    "Kochi",
			State:   "Kerala",
			ZipCode: "98101",
		},
		Role: Department{
			Name:    "Software Engineering",
			Manager: "Sarah",
			Floor:   4,
		},
	}

	jsonData, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}