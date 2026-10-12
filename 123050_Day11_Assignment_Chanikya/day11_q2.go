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

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		{ID: 1, Name: "John", Role: "Developer"},
		{ID: 2, Name: "Alice", Role: "Manager"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	http.HandleFunc("/employee", employeeHandler)

	http.ListenAndServe(":8080", nil)
}
