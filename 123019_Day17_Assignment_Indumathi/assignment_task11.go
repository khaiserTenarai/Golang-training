package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/graphql-go/graphql"
)

type Employee struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
}

// In-Memory Database
var employeesDB = []Employee{
	{ID: "e1", Name: "Vimal", Salary: 60000},
	{ID: "e2", Name: "Aditya", Salary: 50000},
	{ID: "e3", Name: "Adi", Salary: 80000},
}

// Employee Type
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.String},
		"name":   &graphql.Field{Type: graphql.String},
		"salary": &graphql.Field{Type: graphql.Float},
	},
})

// Root Query with Sorting Logic
var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"employees": &graphql.Field{
			Type: graphql.NewList(EmployeeType),
			Args: graphql.FieldConfigArgument{
				"sortBy":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "id"},
				"sortOrder": &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "asc"},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				sortBy := strings.ToLower(p.Args["sortBy"].(string))
				sortOrder := strings.ToLower(p.Args["sortOrder"].(string))

				// Make a copy to avoid mutating the original slice
				sorted := make([]Employee, len(employeesDB))
				copy(sorted, employeesDB)

				// Sort slice based on field and direction
				sort.Slice(sorted, func(i, j int) bool {
					var less bool
					switch sortBy {
					case "name":
						less = sorted[i].Name < sorted[j].Name
					case "salary":
						less = sorted[i].Salary < sorted[j].Salary
					default:
						less = sorted[i].ID < sorted[j].ID
					}

					if sortOrder == "desc" {
						return !less
					}
					return less
				})

				return sorted, nil
			},
		},
	},
})

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: RootQuery})
	if err != nil {
		panic(err)
	}

	// Fetch employees sorted by salary in descending order
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `{ employees(sortBy: "salary", sortOrder: "desc") { id name salary } }`,
	})

	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}