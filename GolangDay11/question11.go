package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var employees = []Employee{
	{ID: 1, Name: "John Doe", Department: "Engineering"},
	{ID: 2, Name: "Jane Smith", Department: "Marketing"},
	{ID: 3, Name: "Alice Johnson", Department: "Engineering"},
	{ID: 4, Name: "Bob Brown", Department: "Sales"},
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	deptFilter := query.Get("department")
	limitStr := query.Get("limit")

	var result []Employee
	for _, emp := range employees {
		if deptFilter == "" || strings.EqualFold(emp.Department, deptFilter) {
			result = append(result, emp)
		}
	}

	if limitStr != "" {
		if limit, err := strconv.Atoi(limitStr); err == nil && limit >= 0 && limit < len(result) {
			result = result[:limit]
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /employees", getEmployees)

	http.ListenAndServe(":8080", mux)
}