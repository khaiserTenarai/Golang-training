package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/graphql-go/graphql"
)

// User Context Struct
type User struct {
	ID   string
	Role string
}

// Authentication Middleware to validate Bearer token and set user in Context
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized: Bearer token missing", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate token (mock validation)
		if token != "valid-auth-token" {
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		user := User{ID: "u101", Role: "ADMIN"}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// Protected Employee Type
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// Protected Query Root
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"me": &graphql.Field{
			Type: EmployeeType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Access user from GraphQL execution context
				user, ok := p.Context.Value("user").(User)
				if !ok {
					return nil, fmt.Errorf("unauthenticated request")
				}
				return map[string]string{"id": user.ID, "name": "Authenticated User"}, nil
			},
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: RootQuery})
	if err != nil {
		panic(err)
	}

	// Simulated Request with Bearer Token in Context
	ctx := context.WithValue(context.Background(), "user", User{ID: "u101", Role: "ADMIN"})

	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `{ me { id name } }`,
		Context:       ctx,
	})

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}