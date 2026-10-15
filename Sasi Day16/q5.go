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

type ListResponse struct {
	Data struct {
		Employees []Employee `json:"employees"`
	} `json:"data"`
}

func employeesListHandler(w http.ResponseWriter, r *http.Request) {
	employees := []Employee{
		{ID: "1", Name: "Sasi", Role: "Engineer"},
		{ID: "2", Name: "Kumar", Role: "Developer"},
		{ID: "3", Name: "Rekha", Role: "Manager"},
	}

	resp := ListResponse{}
	resp.Data.Employees = employees

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql", employeesListHandler)
	fmt.Println("Server running at http://localhost:8085/graphql")
	http.ListenAndServe(":8085", nil)
}
