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

type CreateRequest struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type MutationResponse struct {
	Data struct {
		CreateEmployee Employee `json:"createEmployee"`
	} `json:"data"`
}

var employeesDB = []Employee{}

func createEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var req CreateRequest
		json.NewDecoder(r.Body).Decode(&req)

		newEmp := Employee{
			ID:   fmt.Sprintf("%d", len(employeesDB)+1),
			Name: req.Name,
			Role: req.Role,
		}
		employeesDB = append(employeesDB, newEmp)

		resp := MutationResponse{}
		resp.Data.CreateEmployee = newEmp

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
	http.Error(w, "Use POST method", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/graphql/mutation", createEmployeeHandler)
	fmt.Println("Server running at http://localhost:8086/graphql/mutation")
	http.ListenAndServe(":8086", nil)
}
