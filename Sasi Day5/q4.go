package main

import (
	"encoding/json"
	"fmt"
)

type Address2 struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}

type Department2 struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type Employee4 struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Salary     float64    `json:"salary"`
	Address    Address2    `json:"address"`
	Department Department2 `json:"department"`
}

func main() {

	fmt.Println("\n***************************")
	fmt.Println("4. Marshal Employee to JSON.")
	fmt.Println("****************************")

	emp := Employee4{
		ID:     101,
		Name:   "Sasi",
		Salary: 500000,
		Address: Address2{
			City:    "Madurai",
			State:   "Tamilnadu",
			Pincode: "625301",
		},
		Department: Department2{
			Name:    "Development",
			Manager: "Shiva",
		},
	}

	data, err := json.Marshal(emp)
	
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Println(string(data))
}