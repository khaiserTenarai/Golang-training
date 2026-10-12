package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func getEmployee(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	employee := Employee{
		ID:   id,
		Name: "John",
		Role: "Developer",
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(employee)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /employee/{id}", getEmployee)

	http.ListenAndServe(":8080", mux)
}
