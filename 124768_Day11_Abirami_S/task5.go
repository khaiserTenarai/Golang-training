package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Employee details")
}
func updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Only PUT method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var employee Employee2
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Employee updated")
	fmt.Fprintln(w, employee)
}
func main() {
	http.HandleFunc("/employee", employeeHandler)
	http.HandleFunc("/employee/update", updateEmployeeHandler)
	fmt.Println("Server started on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
