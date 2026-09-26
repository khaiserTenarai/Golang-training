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

	// JSON data
	jsonData := `{
		"id": 101,
		"name": "Pallavi",
		"email": "pallavi@gmail.com",
		"age": 22,
		"salary": 29000
	}`

	var employee Employee

	// Unmarshal converts JSON into Go struct
	err := json.Unmarshal([]byte(jsonData), &employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee Details")
	fmt.Println("----------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
}