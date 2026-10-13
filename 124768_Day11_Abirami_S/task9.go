package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee5 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func employeeDetailHandler(w http.ResponseWriter, r *http.Request) {
	employee := Employee5{
		ID:   101,
		Name: "Abirami",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employee)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]string{
		"error": "Route not found",
	})
}

func main() {
	http.HandleFunc("/employee", employeeDetailHandler)
	http.HandleFunc("/", notFoundHandler)

	fmt.Println("Server started on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}
