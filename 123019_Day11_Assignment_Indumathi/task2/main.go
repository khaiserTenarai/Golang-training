package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

var employees = []Employee{
	{ID: 1, Name: "Indu", Position: "Developer"},
	{ID: 2, Name: "Nan", Position: "Designer"},
}

func getEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	http.HandleFunc("/employee", getEmployeesHandler)
	http.ListenAndServe(":8080", nil)
}