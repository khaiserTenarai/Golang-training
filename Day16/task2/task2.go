package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// 2. Struct representing the Employee type
type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type GraphQLResponse struct {
	Data struct {
		SampleEmployee Employee `json:"sampleEmployee"`
	} `json:"data"`
}

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	emp := Employee{ID: "1", Name: "John Doe", Role: "Developer"}
	res := GraphQLResponse{}
	res.Data.SampleEmployee = emp

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/graphql", employeeHandler)
	fmt.Println("Task 2 Server running at http://localhost:8082/graphql")
	http.ListenAndServe(":8082", nil)
}
