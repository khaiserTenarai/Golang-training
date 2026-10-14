package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

// Module 1: Department Schema Module
var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// Module 2: Employee Schema Module
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.String},
		"name":       &graphql.Field{Type: graphql.String},
		"department": &graphql.Field{Type: DepartmentType},
	},
})

func main() {
	fmt.Println("GraphQL Schema Organized into Modules successfully.")
}
