// 5. Unmarshal JSON into Employee.

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

	jsonData := `{"id": 2, "name": "Ravi", "email": "ravi@example.com", "age": 30, "salary": 60000}`

	var emp Employee
	err := json.Unmarshal([]byte(jsonData), &emp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Unmarshaled Employee:")
	fmt.Printf("%+v\n", emp)
	fmt.Println("Name:", emp.Name)
	fmt.Println("Salary:", emp.Salary)
}
