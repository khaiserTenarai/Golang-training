// 4. Marshal Employee to JSON.

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
	emp := Employee{ID: 1, Name: "Anita", Email: "anita@example.com", Age: 28, Salary: 45000}

	data, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Compact JSON:", string(data))

	prettyData, _ := json.MarshalIndent(emp, "", "  ")
	fmt.Println("Pretty JSON:")
	fmt.Println(string(prettyData))
}
