package main

import (
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// In-Memory Database
var employeesDB = []Employee{
	{ID: "e1", Name: "Vimal"},
	{ID: "e2", Name: "Aditya"},
	{ID: "e3", Name: "Adi"},
	{ID: "e4", Name: "Alex"},
	{ID: "e5", Name: "Mason"},
}

// Employee Type
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// Paginated Employee Response Type
var EmployeePageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EmployeePage",
	Fields: graphql.Fields{
		"items": &graphql.Field{Type: graphql.NewList(EmployeeType)},
		"total": &graphql.Field{Type: graphql.Int},
		"page":  &graphql.Field{Type: graphql.Int},
		"size":  &graphql.Field{Type: graphql.Int},
	},
})

// Root Query with Pagination Logic
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"employees": &graphql.Field{
			Type: EmployeePageType,
			Args: graphql.FieldConfigArgument{
				"page": &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 1},
				"size": &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 2},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				page := p.Args["page"].(int)
				size := p.Args["size"].(int)

				total := len(employeesDB)
				if page < 1 {
					page = 1
				}
				if size < 1 {
					size = 2
				}

				start := (page - 1) * size
				if start >= total {
					return map[string]interface{}{
						"items": []Employee{},
						"total": total,
						"page":  page,
						"size":  size,
					}, nil
				}

				end := start + size
				if end > total {
					end = total
				}

				return map[string]interface{}{
					"items": employeesDB[start:end],
					"total": total,
					"page":  page,
					"size":  size,
				}, nil
			},
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: RootQuery})
	if err != nil {
		panic(err)
	}

	// Fetch Page 1 with Size 2
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `{ employees(page: 1, size: 2) { total page size items { id name } } }`,
	})

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}