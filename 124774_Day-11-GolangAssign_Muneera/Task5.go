package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee4 struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var employees = []Employee4{
	{ID: 1, Name: "Muneera", Department: "IT", Salary: 50000},
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {

	// Get ID from URL
	idString := strings.TrimPrefix(r.URL.Path, "/employee/")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Read JSON from request body
	var employee Employee4

	err = json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Find employee and update
	for i := range employees {

		if employees[i].ID == id {

			employee.ID = id
			employees[i] = employee

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employee)

			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {

	http.HandleFunc("/employee/", updateEmployee)

	fmt.Println("Server running on http://localhost:5050")

	err := http.ListenAndServe(":5050", nil)

	if err != nil {
		fmt.Println(err)
	}
}
