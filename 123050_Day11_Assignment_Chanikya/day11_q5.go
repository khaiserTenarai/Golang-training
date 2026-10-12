package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

var employees = []Employee{
	{ID: 1, Name: "John", Role: "Developer"},
	{ID: 2, Name: "Alice", Role: "Manager"},
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/employee/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	var employee Employee

	err = json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i := range employees {
		if employees[i].ID == id {
			employee.ID = id
			employees[i] = employee

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employee)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/employee/", updateEmployee)

	http.ListenAndServe(":8080", nil)
}
