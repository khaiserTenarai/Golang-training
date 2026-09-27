// 3. Add JSON struct tags.

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID         int     `json:"employee_id"`
	FullName   string  `json:"full_name"`
	Department string  `json:"department,omitempty"`
	Salary     float64 `json:"salary"`
}

func main() {
	emp := Employee{
		ID:         102,
		FullName:   "sahcin",
		Department: "IT",
		Salary:     35000.0,
	}

	data, err := json.Marshal(emp)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}
