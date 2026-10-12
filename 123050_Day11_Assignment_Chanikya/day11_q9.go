package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func createEmployee(w http.ResponseWriter, r *http.Request) {

	var employee Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employee)
}

func main() {
	fmt.Println("Server running on :8080")

	http.HandleFunc("/employee", createEmployee)

	http.ListenAndServe(":8080", nil)
}
