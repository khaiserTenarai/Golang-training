package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		employee := Employee{
			ID:   1,
			Name: "John",
			Role: "Developer",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(employee)

	case http.MethodPost:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(
			map[string]string{
				"message": "Employee created",
			},
		)

	case http.MethodPut:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(
			map[string]string{
				"message": "Employee updated",
			},
		)

	case http.MethodDelete:
		w.WriteHeader(http.StatusNoContent)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/employee", employeeHandler)

	http.ListenAndServe(":8080", nil)
}
