package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/graphql-go/graphql"
)

// Mock DB Models
type Department struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Employee1 struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	DepartmentID string `json:"department_id"`
}

// In-Memory Database
var departmentsDB = map[string]Department{
	"d1": {ID: "d1", Name: "Engineering"},
	"d2": {ID: "d2", Name: "HR"},
}

var employeesDB = []Employee1{
	{ID: "e1", Name: "Sasi", DepartmentID: "d1"},
	{ID: "e2", Name: "Rekha", DepartmentID: "d1"},
	{ID: "e3", Name: "Chandru", DepartmentID: "d2"},
}

// Batch Loader function (Simulates a single batched query to avoid N+1 DB calls)
func batchLoadDepartments(deptIDs []string) map[string]Department {
	results := make(map[string]Department)
	for _, id := range deptIDs {
		if dept, exists := departmentsDB[id]; exists {
			results[id] = dept
		}
	}
	return results
}

// Department Type
var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
	},
})

// Employee Type with Batched Nested Resolver
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee1",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.String},
		"name": &graphql.Field{Type: graphql.String},
		"department": &graphql.Field{
			Type: DepartmentType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				emp, ok := p.Source.(Employee1)
				if !ok {
					return nil, nil
				}

				// Batch fetch optimization using context map
				cachedDepts, ok := p.Context.Value("deptCache").(map[string]Department)
				if ok {
					if dept, exists := cachedDepts[emp.DepartmentID]; exists {
						return dept, nil
					}
				}
				return nil, nil
			},
		},
	},
})

func main() {
	// 1. Collect required IDs for batching
	var deptIDs []string
	for _, emp := range employeesDB {
		deptIDs = append(deptIDs, emp.DepartmentID)
	}

	// 2. Perform single batch load query
	batchedDepts := batchLoadDepartments(deptIDs)

	// 3. Inject batched results into request Context
	ctx := context.WithValue(context.Background(), "deptCache", batchedDepts)

	// 4. Define Root Query
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employees": &graphql.Field{
				Type: graphql.NewList(EmployeeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return employeesDB, nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
	if err != nil {
		panic(err)
	}

	// 5. Execute Query using graphql.Do and graphql.Params
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: `{ employees { id name department { id name } } }`,
		Context:       ctx,
	})

	// Pretty print JSON result
	output, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(output))
}