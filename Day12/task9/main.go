package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type CreateEmployeeDTO struct {
	Name string  `json:"name"`
	Age  int     `json:"age"`
	Role string  `json:"role"`
}

func (dto *CreateEmployeeDTO) Validate() error {
	if strings.TrimSpace(dto.Name) == "" {
		return errors.New("name cannot be empty")
	}
	if dto.Age < 18 {
		return errors.New("employee must be at least 18 years old")
	}
	if dto.Role != "Admin" && dto.Role != "Staff" {
		return errors.New("role must be either 'Admin' or 'Staff'")
	}
	return nil
}

func handleCreate(w http.ResponseWriter, r *http.Request) {
	var dto CreateEmployeeDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := dto.Validate(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"validation_error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Validation passed"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /employee", handleCreate)
	http.ListenAndServe(":8080", mux)
}