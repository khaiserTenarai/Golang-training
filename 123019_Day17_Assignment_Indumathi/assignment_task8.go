package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
)

// Input Validation Function
func validateEmployeeInput(name, email string, salary float64) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is required")
	}
	if !strings.Contains(email, "@") {
		return errors.New("valid email is required")
	}
	if salary < 0 {
		return errors.New("salary cannot be negative")
	}
	return nil
}

// Employee Input Object Type
var EmployeeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EmployeeInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":  &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"salary": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
	},
})

// Employee Object Type
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.String},
		"name":   &graphql.Field{Type: graphql.String},
		"email":  &graphql.Field{Type: graphql.String},
		"salary": &graphql.Field{Type: graphql.Float},
	},
})

// Root Mutation with Input Validation
var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createEmployee": &graphql.Field{
			Type: EmployeeType,
			Args: graphql.FieldConfigArgument{
				"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(EmployeeInputType)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				input := p.Args["input"].(map[string]interface{})
				name := input["name"].(string)
				email := input["email"].(string)
				salary := input["salary"].(float64)

				// Execute Input Validation
				if err := validateEmployeeInput(name, email, salary); err != nil {
					return nil, err
				}

				return map[string]interface{}{
					"id":     "e101",
					"name":   name,
					"email":  email,
					"salary": salary,
				}, nil
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

	// Test Case 1: Invalid Input (Negative Salary)
	res1 := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `mutation { createEmployee(input: { name: "Vimal", email: "vimal@gmail.com", salary: -5000 }) { id name } }`,
	})
	out1, _ := json.MarshalIndent(res1, "", "  ")
	fmt.Println("Invalid Input Output:")
	fmt.Println(string(out1))

	// Test Case 2: Valid Input
	res2 := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `mutation { createEmployee(input: { name: "Vimal", email: "vimal@gmail.com", salary: 60000 }) { id name email salary } }`,
	})
	out2, _ := json.MarshalIndent(res2, "", "  ")
	fmt.Println("\nValid Input Output:")
	fmt.Println(string(out2))
}