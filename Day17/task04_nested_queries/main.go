package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DepartmentID string `json:"departmentId"`
}

var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"department": &graphql.Field{
			Type: DepartmentType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				emp := p.Source.(Employee)
				// Fetch nested department
				return Department{ID: emp.DepartmentID, Name: "Engineering"}, nil
			},
		},
	},
})

func main() {
	fmt.Println("Nested Employee -> Department query configuration defined.")
}
