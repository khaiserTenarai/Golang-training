package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee3 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var employee Employee3
		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "Employee created successfully")
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintln(w, "Only POST method is allowed")
}
func main() {
	http.HandleFunc("/employee", employeeHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
