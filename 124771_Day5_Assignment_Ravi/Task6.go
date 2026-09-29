package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}

type Employee struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Address Address `json:"address"`
}

func main() {
	jsonData := `{
		"id": 101,
		"name": "Rajesh",
		"address": {
			"city": "Bangalore",
			"state": "Karnataka",
			"pincode": "111111"
		}
	}`

	var employee Employee

	err := json.Unmarshal([]byte(jsonData), &employee)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(employee)
	fmt.Println(employee.Name)
	fmt.Println(employee.Address.City)
}
