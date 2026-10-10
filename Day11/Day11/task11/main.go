package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Department string `json:"department"`
}

var employees = []Employee{
	{ID: 1, Name: "Ray", Department: "Engineering"},
	{ID: 2, Name: "coco", Department: "Design"},
	{ID: 3, Name: "stella", Department: "Engineering"},
}

func listEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query params (e.g., /employees?department=Engineering)
	deptQuery := r.URL.Query().Get("department")

	var result []Employee
	for _, emp := range employees {
		if deptQuery == "" || strings.EqualFold(emp.Department, deptQuery) {
			result = append(result, emp)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", listEmployeesHandler)
	http.ListenAndServe(":8080", mux)
}