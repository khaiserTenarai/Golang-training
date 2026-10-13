package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee7 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var employees = []Employee7{
	{ID: 101, Name: "Abirami"},
	{ID: 102, Name: "Priya"},
	{ID: 103, Name: "Rahul"},
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	http.HandleFunc("/employees", employeesHandler)

	fmt.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
