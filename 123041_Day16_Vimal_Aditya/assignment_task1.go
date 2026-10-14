package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const typeDefs = `
type Query {
  healthCheck: String!
}
`

type SchemaResponse struct {
	Schema string `json:"schema"`
	Status string `json:"status"`
}

func schemaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(SchemaResponse{
		Schema: typeDefs,
		Status: "GraphQL Schema created successfully!",
	})
}

func main() {
	http.HandleFunc("/graphql/schema", schemaHandler)
	fmt.Println("Server running at http://localhost:8081/graphql/schema")
	http.ListenAndServe(":8081", nil)
}
