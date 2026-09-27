package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string `json:"city"`
	Pincode string `json:"pincode"`
}

type Department struct {
	Name string `json:"name"`
}

type Employee struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	Email      string     `json:"email"`
	Age        int        `json:"age"`
	Salary     float64    `json:"salary"`
	Address    Address    `json:"address"`
	Department Department `json:"department"`
}

func main() {
	
	jsonText := `{
		"id": 2,
		"name": "Meera",
		"email": "meera@example.com",
		"age": 29,
		"salary": 70000,
		"address": {"city": "Mysuru", "pincode": "570001"},
		"department": {"name": "HR"}
	}`

	var e Employee
	err := json.Unmarshal([]byte(jsonText), &e) 
	if err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}

	fmt.Println("JSON unmarshaled back into an Employee struct:")
	fmt.Printf("%+v\n", e)
	fmt.Println("Name field:", e.Name)
	fmt.Println("City field:", e.Address.City)
}
