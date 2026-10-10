package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type EmployeeRequest struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Salary   float64 `json:"salary"`
}

func (e *EmployeeRequest) Validate() error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(e.Email, "@") {
		return errors.New("invalid email address")
	}
	if e.Salary <= 0 {
		return errors.New("salary must be greater than zero")
	}
	return nil
}

func createValidatedEmployee(w http.ResponseWriter, r *http.Request) {
	var req EmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	if err := req.Validate(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity) // 422
		json.NewEncoder(w).Encode(map[string]string{"validation_error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee validation passed and created successfully!"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /employee", createValidatedEmployee)
	http.ListenAndServe(":8080", mux)
}