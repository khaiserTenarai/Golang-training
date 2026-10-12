package main

import (
	"fmt"
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

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	idString := strings.TrimPrefix(r.URL.Path, "/employee/")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, employee := range employees {
		if employee.ID == id {
			employees = append(employees[:i], employees[i+1:]...)

			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/employee/", deleteEmployee)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
