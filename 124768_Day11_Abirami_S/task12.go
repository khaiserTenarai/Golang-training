package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Employee8 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var employees = []Employee8{
	{ID: 101, Name: "Abirami"},
	{ID: 102, Name: "Priya"},
	{ID: 103, Name: "Rahul"},
}

func employeeByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idText := strings.TrimPrefix(r.URL.Path, "/employees/")

	id, err := strconv.Atoi(idText)
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	for _, employee := range employees {
		if employee.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(employee)
			return
		}
	}

	http.Error(w, "Employee not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/employees/", employeeByIDHandler)

	fmt.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
