package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

var nextID = 3
var mockDB = make(map[int]Employee)

func createEmployee(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	emp.ID = nextID
	nextID++
	mockDB[emp.ID] = emp

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	http.HandleFunc("/employee", createEmployee)
	http.ListenAndServe(":8080", nil)
}