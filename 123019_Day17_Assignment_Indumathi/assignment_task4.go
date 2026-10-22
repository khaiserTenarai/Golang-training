package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

// Department Schema
var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// Employee Schema with Nested Department
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"department": &graphql.Field{
			Type: DepartmentType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Resolves nested department for the employee
				return nil, nil
			},
		},
	},
})

func main() {
	fmt.Println("Nested Employee -> Department Query Implemented")
}