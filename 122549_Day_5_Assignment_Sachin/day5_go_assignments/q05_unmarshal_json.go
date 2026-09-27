// 5. Unmarshal JSON into Employee.

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Role   string  `json:"role"`
	Salary float64 `json:"salary"`
}

func main() {
	rawJSON := `{"id": 101, "name": "sachin", "role": "eng", "salary": 72000}`

	var emp Employee
	err := json.Unmarshal([]byte(rawJSON), &emp)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	fmt.Println("Unmarshalled Struct:")
	fmt.Printf("ID: %d, Name: %s, Role: %s, Salary: %.2f\n", emp.ID, emp.Name, emp.Role, emp.Salary)
}
