package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

func main() {
	gokul := Employee{
		ID:     1,
		Name:   "Gokul",
		Email:  "gokul@example.com",
		Salary: 65000,
	}

	data, err := json.Marshal(gokul)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println("Compact JSON:", string(data))

	// MarshalIndent gives a nicer, readable output
	prettyData, _ := json.MarshalIndent(gokul, "", "  ")
	fmt.Println("\nPretty JSON:")
	fmt.Println(string(prettyData))
}
