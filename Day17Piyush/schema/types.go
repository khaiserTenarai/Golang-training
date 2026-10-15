// Task 1: Organize GraphQL schema into modules — Type definitions

package schema

import (
	"github.com/graphql-go/graphql"
)

// Department type
var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int},
		"name":     &graphql.Field{Type: graphql.String},
		"location": &graphql.Field{Type: graphql.String},
	},
})

// Employee type — Task 4: nested Department field added in InitTypes()
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.Int},
		"name":          &graphql.Field{Type: graphql.String},
		"email":         &graphql.Field{Type: graphql.String},
		"age":           &graphql.Field{Type: graphql.Int},
		"department_id": &graphql.Field{Type: graphql.Int},
		"salary":        &graphql.Field{Type: graphql.Float},
		"role":          &graphql.Field{Type: graphql.String},
		"created_at":    &graphql.Field{Type: graphql.String},
	},
})

// Paginated employee list (Task 9)
var EmployeeListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EmployeeList",
	Fields: graphql.Fields{
		"employees":   &graphql.Field{Type: graphql.NewList(EmployeeType)},
		"total":       &graphql.Field{Type: graphql.Int},
		"page":        &graphql.Field{Type: graphql.Int},
		"page_size":   &graphql.Field{Type: graphql.Int},
		"total_pages": &graphql.Field{Type: graphql.Int},
	},
})

// Auth response
var AuthResponseType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuthResponse",
	Fields: graphql.Fields{
		"token":    &graphql.Field{Type: graphql.String},
		"user_id":  &graphql.Field{Type: graphql.Int},
		"username": &graphql.Field{Type: graphql.String},
		"role":     &graphql.Field{Type: graphql.String},
	},
})

// Employee input for mutations (Task 8: validation happens in service)
var EmployeeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "EmployeeInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"age":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"department_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"salary":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
		"role":          &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})
