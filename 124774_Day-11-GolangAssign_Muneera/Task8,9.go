package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee5 struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// Question 8: Return JSON response
func getEmployee(w http.ResponseWriter, r *http.Request) {

	employee := Employee5{
		ID:         1,
		Name:       "Muneera",
		Department: "IT",
		Salary:     50000,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employee)
}

// Question 9: Accept JSON request body
func createEmployee(w http.ResponseWriter, r *http.Request) {

	var employee Employee5

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employee)
}

func main() {

	http.HandleFunc("/employee", getEmployee)

	http.HandleFunc("/employee/create", createEmployee)

	fmt.Println("Server running on http://localhost:1010")

	err := http.ListenAndServe(":1010", nil)

	if err != nil {
		fmt.Println(err)
	}
}
