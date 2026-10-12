package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

func main() {

	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			employees := []Employee{
				{ID: 1, Name: "Sanskriti", Email: "sanskriti@gmail.com", Salary: 50000},
				{ID: 2, Name: "Rahul", Email: "rahul@gmail.com", Salary: 60000},
			}

			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(employees)
		}
	})

	http.ListenAndServe(":8080", nil)
}