package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {

	role := r.URL.Query().Get("role")

	employees := []Employee{
		{ID: 1, Name: "John", Role: "Developer"},
		{ID: 2, Name: "Alice", Role: "Manager"},
		{ID: 3, Name: "David", Role: "Developer"},
	}

	var result []Employee

	for _, employee := range employees {

		if role == "" || employee.Role == role {
			result = append(result, employee)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}

func main() {

	http.HandleFunc("/employee", getEmployees)

	http.ListenAndServe(":8080", nil)
}
