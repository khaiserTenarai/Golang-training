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

// Schema Setup for Testing
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"employee": &graphql.Field{
			Type: EmployeeType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return map[string]string{"id": "e101", "name": "Vimal"}, nil
			},
		},
	},
})

// Unit Test for GraphQL HTTP Endpoint
func TestGraphQLEndpoint(t *testing.T) {
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: RootQuery})
	h := handler.New(&handler.Config{Schema: &schema})

	// Simulate POST HTTP request with GraphQL query
	query := `{"query": "{ employee { id name } }"}`
	req := httptest.NewRequest(http.MethodPost, "/graphql", strings.NewReader(query))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	// Verify HTTP Status Code
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200 OK, got %d", rec.Code)
	}

	fmt.Println("GraphQL Query Test Passed successfully!")
}

func main() {
	// Runs the test simulation inside main for verification
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: RootQuery})
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `{ employee { id name } }`,
	})

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println("Test Query Execution Output:")
	fmt.Println(string(output))
}