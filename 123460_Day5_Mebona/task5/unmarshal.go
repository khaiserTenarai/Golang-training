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
	jsonData := `{
		"id": 101,
		"name": "John",
		"email": "john@gmail.com",
		"age": 28,
		"salary": 50000
	}`

	var employee Employee

	err := json.Unmarshal([]byte(jsonData), &employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
}