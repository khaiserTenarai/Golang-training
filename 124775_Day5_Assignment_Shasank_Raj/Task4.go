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
	employee := Employee{
		ID:   101,
		Name: "jijuuu",
		Address: Address{
			City:    "Bangalore",
			State:   "Karnataka",
			Pincode: "836203",
		},
	}

	data, _ := json.Marshal(employee)

	fmt.Println(string(data))
}
