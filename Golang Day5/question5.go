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
	jsonInput := `{
		"id": 205,
		"name": "Lakshmi Shibu",
		"location": {
			"street": "456 MG road",
			"city": "Kochi",
			"state": "Kerala",
			"zip": "73301"
		},
		"role": {
			"department": "Data Science",
			"manager": "Rahul",
			"floor": 2
		}
	}`

	var emp Employee

	err := json.Unmarshal([]byte(jsonInput), &emp)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(emp.FullName)
	fmt.Println(emp.Location.City)
	fmt.Println(emp.Role.Name)
}