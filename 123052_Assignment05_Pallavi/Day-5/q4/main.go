package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Age    int     `json:"age"`
	Salary float64 `json:"salary"`
}

func main() {

	employee := Employee{
		ID:     101,
		Name:   "Pallavi",
		Email:  "pallavi@gmail.com",
		Age:    22,
		Salary: 29000,
	}

	// Marshal converts Go struct into JSON
	data, err := json.Marshal(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("JSON Data:")
	fmt.Println(string(data))
}