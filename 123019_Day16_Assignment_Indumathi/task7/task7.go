package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type UpdateRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type UpdateResponse struct {
	Data struct {
		UpdateEmployee *Employee `json:"updateEmployee"`
	} `json:"data"`
}

var employeesDB = []Employee{
	{ID: "1", Name: "Alex", Role: "Junior Dev"},
}

// 7. Mutation for updating an existing employee
func updateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateRequest
	json.NewDecoder(r.Body).Decode(&req)

	var updated *Employee
	for i, emp := range employeesDB {
		if emp.ID == req.ID {
			if req.Name != "" {
				employeesDB[i].Name = req.Name
			}
			if req.Role != "" {
				employeesDB[i].Role = req.Role
			}
			updated = &employeesDB[i]
			break
		}
	}

	resp := UpdateResponse{}
	resp.Data.UpdateEmployee = updated

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql/update", updateEmployeeHandler)
	fmt.Println("Task 7 Server running at http://localhost:8087/graphql/update")
	http.ListenAndServe(":8087", nil)
}
