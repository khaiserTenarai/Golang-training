package main

import (
	"encoding/json"
	"net/http"
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

func employeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/employee/")

	var employee Employee

	for _, e := range employees {
		if string(rune(e.ID+'0')) == id {
			employee = e
			break
		}
	}

	if employee.ID == 0 {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func main() {
	http.HandleFunc("/employee/", employeeDetailHandler)

	http.ListenAndServe(":8080", nil)
}
