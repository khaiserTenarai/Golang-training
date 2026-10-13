package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee9 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var employees = []Employee9{{101, "Abirami"}, {102, "Priya"}, {103, "Rahul"}}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(employees)
	case http.MethodPost:
		var e Employee9
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		for _, old := range employees {
			if old.ID == e.ID {
				http.Error(w, "Employee ID already exists", http.StatusConflict)
				return
			}
		}
		employees = append(employees, e)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(e)
	default:
		w.Header().Set("Allow", "GET, POST")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func employeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/employees/"))
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	for _, e := range employees {
		if e.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(e)
			return
		}
	}
	http.Error(w, "Employee not found", http.StatusNotFound)
}
func main() {
	http.HandleFunc("/employees", employeesHandler)
	http.HandleFunc("/employees/", employeeByIDHandler)
	fmt.Println("Server started on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
