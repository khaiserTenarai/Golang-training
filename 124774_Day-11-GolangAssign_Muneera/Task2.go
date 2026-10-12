package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {

	employees := []Employee{
		{
			ID:         1,
			Name:       "Muneera",
			Department: "IT",
		},
		{
			ID:         2,
			Name:       "Rahul",
			Department: "HR",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employees)
}

func main() {

	http.HandleFunc("/employee", getEmployees)

	fmt.Println("Server running on http://localhost:9090")

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		fmt.Println(err)
	}
}

//http://localhost:8080/employee
