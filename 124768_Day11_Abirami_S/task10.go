package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee6 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		employee := Employee6{
			ID:   101,
			Name: "Abirami",
		}

		json.NewEncoder(w).Encode(employee)

	case http.MethodPost:
		var employee Employee6

		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, `{"error":"Invalid JSON"}`, http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(employee)

	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/employee", employeeHandler)

	fmt.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
