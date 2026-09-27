package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

func main() {
	jsonData := `{"id": 2, "name": "Gokul", "email": "gokul@example.com", "salary": 65000}`

	var emp Employee
	err := json.Unmarshal([]byte(jsonData), &emp)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	fmt.Printf("Parsed struct: %+v\n", emp)
	fmt.Println("Name from JSON:", emp.Name)
	fmt.Println("Salary from JSON:", emp.Salary)
}
