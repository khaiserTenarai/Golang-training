// 4. Marshal Employee to JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func main() {
	emp := Employee{
		ID:       501,
		Name:     "gokul",
		Position: "Developer",
		Salary:   80000,
	}

	jsonBytes, err := json.MarshalIndent(emp, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println("JSON Output:")
	fmt.Println(string(jsonBytes))
}
