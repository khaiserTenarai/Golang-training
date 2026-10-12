package main

import (
	"encoding/json"
	"fmt"
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

		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var employee Employee

		err := json.NewDecoder(r.Body).Decode(&employee)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		employee.ID = 1

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(employee)
	})
	fmt.Println("http://localhost:8080/employee")
	http.ListenAndServe(":8080", nil)
}