package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}

type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Age         int        `json:"age"`
	Salary      float64    `json:"salary"`
	Phone       string     `json:"phone"`
	Position    string     `json:"position"`
	JoiningDate string     `json:"joining_date"`
	Address     Address    `json:"address"`
	Department  Department `json:"department"`
}

func main() {

	// JSON data
	jsonData := `{
		"id": 101,
		"name": "Muneera",
		"email": "muneera@gmail.com",
		"age": 21,
		"salary": 20000,
		"phone": "9876543210",
		"position": "Software Developer",
		"joining_date": "16-09-2026",
		"address": {
			"city": "Bangalore",
			"state": "Karnataka",
			"pincode": "560037"
		},
		"department": {
			"id": 10,
			"name": "IT"
		}
	}`

	var employee Employee

	// Convert JSON into Employee struct
	err := json.Unmarshal([]byte(jsonData), &employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("------Employee Details------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Phone:", employee.Phone)
	fmt.Println("Position:", employee.Position)
	fmt.Println("Joining Date:", employee.JoiningDate)

	fmt.Println("\n-----Address------")
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Pincode:", employee.Address.Pincode)

	fmt.Println("\n-----Department-----")
	fmt.Println("Department ID:", employee.Department.ID)
	fmt.Println("Department Name:", employee.Department.Name)
}
