package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

var employees = map[int]Employee{
	1: {ID: 1, Name: "Vimal", Position: "Manager"},
	2: {ID: 2, Name: "Aditya", Position: "Data Engineer"},
}

func employeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	emp, exists := employees[id]
	if !exists {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employee/{id}", employeeDetailHandler)
	http.ListenAndServe(":8080", mux)
}