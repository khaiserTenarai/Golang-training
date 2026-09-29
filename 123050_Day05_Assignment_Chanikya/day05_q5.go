package main

import (
	"encoding/json"
	"fmt"
)

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

	jsonData := `{
		"id": 101,
		"name": "Ram",
		"address": {
			"city": "Bangalore",
			"state": "Karnataka"
		},
		"department": {
			"name": "IT"
		}
	}`

	var employee Employee

	err := json.Unmarshal([]byte(jsonData), &employee)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(employee.Name)
	fmt.Println(employee.Address.City)
	fmt.Println(employee.Department.Name)
}
