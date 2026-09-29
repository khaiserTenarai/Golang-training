// 3. Add JSON struct tags.

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
	emp := Employee{ID: 1, Name: "Anita Sharma", Email: "anita@example.com", Age: 28, Salary: 45000}

	data, _ := json.Marshal(emp)
	fmt.Println(string(data))

}
