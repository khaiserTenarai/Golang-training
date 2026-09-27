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
		Name:   "John",
		Email:  "john@gmail.com",
		Age:    28,
		Salary: 50000,
	}

	jsonData, err := json.Marshal(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}