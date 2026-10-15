package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GraphQLPayload struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type VariableResponse struct {
	Data struct {
		Employee map[string]interface{} `json:"employee"`
	} `json:"data"`
}

func variablesHandler(w http.ResponseWriter, r *http.Request) {
	var payload GraphQLPayload
	json.NewDecoder(r.Body).Decode(&payload)

	empID := payload.Variables["empId"]

	resp := VariableResponse{}
	resp.Data.Employee = map[string]interface{}{
		"id":     empID,
		"name":   fmt.Sprintf("Employee_%v", empID),
		"salary": 75000.00,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/graphql/variables", variablesHandler)
	fmt.Println("Server running at http://localhost:8092/graphql/variables")
	http.ListenAndServe(":8092", nil)
}
