package main

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type Employee struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}

func validateEmployee(emp Employee) string {
	if emp.Name == "" {
		return "Name is required"
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(emp.Email) {
		return "Invalid email format"
	}

	if emp.Position == "" {
		return "Position is required"
	}

	if emp.Salary <= 0 {
		return "Salary must be greater than zero"
	}

	return ""
}

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if validationErr := validateEmployee(emp); validationErr != "" {
		http.Error(w, validationErr, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee validated and created successfully"})
}

func main() {
	http.HandleFunc("/employee", createEmployeeHandler)
	http.ListenAndServe(":8080", nil)
}