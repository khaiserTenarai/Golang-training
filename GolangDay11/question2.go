package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	employee := Employee{
		ID:         1,
		Name:       "Lakshmi",
		Department: "Engineering",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func main() {
	http.HandleFunc("/employee", employeeHandler)
	http.ListenAndServe(":8080", nil)
}