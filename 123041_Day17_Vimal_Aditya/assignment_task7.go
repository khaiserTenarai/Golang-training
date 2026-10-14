package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

// User Context Struct
type User struct {
	ID   string
	Role string
}

// Authorization Helper
func requireAdmin(ctx context.Context) error {
	user, ok := ctx.Value("user").(User)
	if !ok {
		return fmt.Errorf("unauthenticated")
	}
	if user.Role != "ADMIN" {
		return fmt.Errorf("forbidden: admin authorization required")
	}
	return nil
}

// Root Mutation with Admin Authorization Check
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"deleteEmployee": &graphql.Field{
			Type: graphql.Boolean,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Check if user is authorized as ADMIN
				if err := requireAdmin(p.Context); err != nil {
					return false, err
				}
				return true, nil // Authorized: execute deletion
			},
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"health": &graphql.Field{
					Type: graphql.String,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return "OK", nil
					},
				},
			},
		}),
		Mutation: RootMutation,
	})
	if err != nil {
		panic(err)
	}

	// Test Case 1: USER Role (Should fail authorization)
	userCtx := context.WithValue(context.Background(), "user", User{ID: "u102", Role: "USER"})
	res1 := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `mutation { deleteEmployee(id: "e101") }`,
		Context:       userCtx,
	})
	out1, _ := json.MarshalIndent(res1, "", "  ")
	fmt.Println("USER Role Output:")
	fmt.Println(string(out1))

	// Test Case 2: ADMIN Role (Should succeed)
	adminCtx := context.WithValue(context.Background(), "user", User{ID: "u101", Role: "ADMIN"})
	res2 := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `mutation { deleteEmployee(id: "e101") }`,
		Context:       adminCtx,
	})
	out2, _ := json.MarshalIndent(res2, "", "  ")
	fmt.Println("\nADMIN Role Output:")
	fmt.Println(string(out2))
}