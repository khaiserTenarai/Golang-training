package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Employee struct {
	Name string `json:"name"`
	Dept string `json:"dept"`
}

var list = []Employee{
	{Name: "Indu", Dept: "Engineering"},
	{Name: "Nan", Dept: "HR"},
	{Name: "Aro", Dept: "Engineering"},
}

func filterHandler(w http.ResponseWriter, r *http.Request) {
	deptFilter := r.URL.Query().Get("dept")

	var filtered []Employee
	for _, emp := range list {
		if deptFilter == "" || strings.EqualFold(emp.Dept, deptFilter) {
			filtered = append(filtered, emp)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filtered)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", filterHandler)
	http.ListenAndServe(":8080", mux)
}