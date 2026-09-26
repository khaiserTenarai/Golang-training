package main

import (
	"encoding/json"
	"fmt"
)

// Employee contains JSON struct tags
type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

func main() {

	employee := Employee{
		ID:         101,
		Name:       "Pallavi",
		Email:      "pallavi@gmail.com",
		Age:        22,
		Department: "IT",
		Salary:     29000,
	}

	// Convert Employee struct to JSON
	data, _ := json.Marshal(employee)

	fmt.Println(string(data))
}