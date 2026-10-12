package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Employee struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func validateEmployee(employee Employee) bool {

	if strings.TrimSpace(employee.Name) == "" {
		return false
	}

	if strings.TrimSpace(employee.Email) == "" {
		return false
	}

	if strings.TrimSpace(employee.Role) == "" {
		return false
	}

	return true
}

func createEmployee(w http.ResponseWriter, r *http.Request) {

	var employee Employee

	err := json.NewDecoder(r.Body).Decode(&employee)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if !validateEmployee(employee) {
		http.Error(
			w,
			"Name, email and role are required",
			http.StatusBadRequest,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(employee)
}

func main() {

	http.HandleFunc("/employee", createEmployee)

	http.ListenAndServe(":8080", nil)
}
