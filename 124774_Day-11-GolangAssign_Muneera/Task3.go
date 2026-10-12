package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee1 struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var employees = []Employee1{
	{
		ID:         1,
		Name:       "Muneera",
		Department: "IT",
		Salary:     50000,
	},
	{
		ID:         2,
		Name:       "Rahul",
		Department: "HR",
		Salary:     45000,
	},
}

func employeeDetail(w http.ResponseWriter, r *http.Request) {

	// Extract ID from the URL
	idString := strings.TrimPrefix(r.URL.Path, "/employee/")

	// Convert string ID to integer
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Find employee by ID
	for _, employee := range employees {
		if employee.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employee)
			return
		}
	}

	// Employee not found
	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/employee/", employeeDetail)
	fmt.Println("Server running on http://localhost:7070")
	err := http.ListenAndServe(":7070", nil)

	if err != nil {
		fmt.Println(err)
	}
}

//http://localhost:8080/employee/1
