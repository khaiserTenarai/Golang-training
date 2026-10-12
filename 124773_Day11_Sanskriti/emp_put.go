package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"fmt"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Salary float64 `json:"salary"`
}

func main() {

	employee := Employee{
		ID:     1,
		Name:   "Sanskriti",
		Email:  "sanskriti@gmail.com",
		Salary: 50000,
	}

	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPut {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		idString := strings.TrimPrefix(r.URL.Path, "/employee/")

		id, err := strconv.Atoi(idString)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		if id != employee.ID {
			http.Error(w, "Employee not found", http.StatusNotFound)
			return
		}

		err = json.NewDecoder(r.Body).Decode(&employee)

		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		employee.ID = id

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(employee)
	})
	fmt.Println("http://localhost:8080/employee")
	http.ListenAndServe(":8080", nil)
}