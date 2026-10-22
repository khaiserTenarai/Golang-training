package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Employee struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

// 10. Go struct matching a GraphQL Input Type
type EmployeeInput struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}

type InputTypeResponse struct {
	Data struct {
		AddEmployee Employee `json:"addEmployee"`
	} `json:"data"`
}

func inputTypeHandler(w http.ResponseWriter, r *http.Request) {
	var input EmployeeInput
	json.NewDecoder(r.Body).Decode(&input)

	newEmp := Employee{
		ID:         "emp-101",
		Name:       input.Name,
		Department: input.Department,
	}

	resp := InputTypeResponse{}
	resp.Data.AddEmployee = newEmp

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql/input", inputTypeHandler)
	fmt.Println("Task 10 Server running at http://localhost:8090/graphql/input")
	http.ListenAndServe(":8090", nil)
}
