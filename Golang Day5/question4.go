package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip"`
}

type Department struct {
	Name    string `json:"department"`
	Manager string `json:"manager"`
	Floor   int    `json:"floor"`
}

type Employee struct {
	ID       int        `json:"id"`
	FullName string     `json:"name"`
	Location Address    `json:"location"`
	Role     Department `json:"role"`
}

func main() {
	emp := Employee{
		ID:       205,
		FullName: "Lakshmi Shibu",
		Location: Address{
			Street:  "456 Data Drive",
			City:    "Kochi",
			State:   "Kerala",
			ZipCode: "73301",
		},
		Role: Department{
			Name:    "Data Science",
			Manager: "Rahul",
			Floor:   2,
		},
	}

	jsonData, err := json.Marshal(emp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}