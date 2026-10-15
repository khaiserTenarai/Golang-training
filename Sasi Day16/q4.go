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

type QueryResponse struct {
	Data struct {
		Employee Employee `json:"employee"`
	} `json:"data"`
}

func queryEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	emp := Employee{ID: "emp-1", Name: "Sasi", Role: "Developer"}

	resp := QueryResponse{}
	resp.Data.Employee = emp

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql", queryEmployeeHandler)
	fmt.Println("Server running at http://localhost:8084/graphql")
	http.ListenAndServe(":8084", nil)
}
