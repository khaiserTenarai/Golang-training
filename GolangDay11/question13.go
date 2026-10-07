package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
	Age        int    `json:"age"`
}

func (e *Employee) Validate() error {
	if strings.TrimSpace(e.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(e.Department) == "" {
		return errors.New("department is required")
	}
	if e.Age <= 0 || e.Age > 120 {
		return errors.New("age must be between 1 and 120")
	}
	return nil
}

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var emp Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if err := emp.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(emp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /employee", createEmployeeHandler)

	http.ListenAndServe(":8080", mux)
}