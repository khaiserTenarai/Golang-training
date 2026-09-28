package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode int    `json:"pincode"`
}

type Department struct {
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Age        int        `json:"age"`
	Salary     float64    `json:"salary"`
	Phone      string     `json:"phone"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
}

func main() {

	jsonData := `{
		"id": 101,
		"name": "Swathi",
		"email": "swathi@gmail.com",
		"age": 25,
		"salary": 50000,
		"phone": "9876543210",
		"address": {
			"street": "MG Road",
			"city": "Bangalore",
			"state": "Karnataka",
			"pincode": 560001
		},
		"department": {
			"name": "IT",
			"manager": "Rahul"
		}
	}`

	var employee Employee

	err := json.Unmarshal([]byte(jsonData), &employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Employee Email:", employee.Email)
	fmt.Println("Employee Age:", employee.Age)
	fmt.Println("Employee Salary:", employee.Salary)
	fmt.Println("Employee Phone:", employee.Phone)

	fmt.Println("\nAddress:")
	fmt.Println("Street:", employee.Address.Street)
	fmt.Println("City:", employee.Address.City)
	fmt.Println("State:", employee.Address.State)
	fmt.Println("Pincode:", employee.Address.Pincode)

	fmt.Println("\nDepartment:")
	fmt.Println("Department Name:", employee.Department.Name)
	fmt.Println("Manager:", employee.Department.Manager)
}