package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

var employees = []Employee{
	{ID: 1, Name: "Anita Rao", Email: "anita@example.com", Department: "Engineering", Salary: 85000},
	{ID: 2, Name: "Rahul Verma", Email: "rahul@example.com", Department: "Finance", Salary: 62000},
	{ID: 3, Name: "Sara Khan", Email: "sara@example.com", Department: "Engineering", Salary: 91000},
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/employee", employeesHandler)

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
