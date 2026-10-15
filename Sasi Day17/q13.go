package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Employee Type
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// Root Mutation
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createEmployee": &graphql.Field{
			Type: EmployeeType,
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				name := p.Args["name"].(string)
				return map[string]string{"id": "e102", "name": name}, nil
			},
		},
	},
})

// Unit Test for GraphQL Mutation HTTP Endpoint
func TestCreateEmployeeMutation(t *testing.T) {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"health": &graphql.Field{Type: graphql.String, Resolve: func(p graphql.ResolveParams) (interface{}, error) { return "OK", nil }},
			},
		}),
		Mutation: RootMutation,
	})

	h := handler.New(&handler.Config{Schema: &schema})

	// Simulate POST HTTP request with GraphQL mutation
	mutationBody := `{"query": "mutation { createEmployee(name: \"Vimal\") { id name } }"}`
	req := httptest.NewRequest(http.MethodPost, "/graphql", strings.NewReader(mutationBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	// Verify HTTP Status Code
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", rec.Code)
	}

	fmt.Println("GraphQL Mutation Test Passed successfully!")
}

func main() {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"health": &graphql.Field{Type: graphql.String, Resolve: func(p graphql.ResolveParams) (interface{}, error) { return "OK", nil }},
			},
		}),
		Mutation: RootMutation,
	})

	// Test Mutation Execution Output
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `mutation { createEmployee(name: "Sasi") { id name } }`,
	})

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println("Test Mutation Execution Output:")
	fmt.Println(string(output))
}