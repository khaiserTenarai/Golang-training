package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/graphql-go/graphql"
)

type Employee struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-Memory Database
var employeesDB = []Employee{
	{ID: "e1", Name: "Sasi", Email: "sasi@gmail.com"},
	{ID: "e2", Name: "Rekha", Email: "rekha@yahoo.com"},
	{ID: "e3", Name: "Vishwa", Email: "vishwa@gmail.com"},
}

// Employee Type
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String},
		"name":  &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
	},
})

// Root Query with Filtering Logic
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"employees": &graphql.Field{
			Type: graphql.NewList(EmployeeType),
			Args: graphql.FieldConfigArgument{
				"search": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: ""},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				search := strings.ToLower(p.Args["search"].(string))

				// If no search filter provided, return all records
				if search == "" {
					return employeesDB, nil
				}

				// Filter records by matching name or email
				var filtered []Employee
				for _, emp := range employeesDB {
					if strings.Contains(strings.ToLower(emp.Name), search) ||
						strings.Contains(strings.ToLower(emp.Email), search) {
						filtered = append(filtered, emp)
					}
				}

				return filtered, nil
			},
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: RootQuery})
	if err != nil {
		panic(err)
	}

	// Filter Employees matching search string "gmail"
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `{ employees(search: "gmail") { id name email } }`,
	})

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}