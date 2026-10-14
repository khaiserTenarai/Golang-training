package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// ============================================================================
// DATA MODELS
// ============================================================================

// Department represents a department entity (Requirement #3)
type Department struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

// Employee represents an employee entity (Requirement #2)
type Employee struct {
	ID           int32   `json:"id"`
	Name         string  `json:"name"`
	Email        string  `json:"email"`
	Salary       float64 `json:"salary"`
	IsActive     bool    `json:"isActive"`
	DepartmentID string  `json:"departmentId"`
}

// In-memory data source
var departments = []Department{
	{ID: "d1", Name: "Engineering", Location: "Building A"},
	{ID: "d2", Name: "Human Resources", Location: "Building B"},
}

var employees = []Employee{
	{ID: 1, Name: "Pallavi", Email: "pallavi@gmail.com", Salary: 85000.0, IsActive: true, DepartmentID: "d1"},
	{ID: 2, Name: "Amit Sharma", Email: "amit@gmail.com", Salary: 65000.0, IsActive: true, DepartmentID: "d2"},
}

// ============================================================================
// GRAPHQL TYPES & SCALARS (Requirements #1, #2, #3, #9, #10)
// ============================================================================

// 3. Department Type
var departmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.String},
		"name":     &graphql.Field{Type: graphql.String},
		"location": &graphql.Field{Type: graphql.String},
	},
})

// 2 & 9. Employee Type using Scalar Types (ID, String, Float, Boolean)
var employeeType *graphql.Object

func initTypes() {
	employeeType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"id":       &graphql.Field{Type: graphql.ID},
			"name":     &graphql.Field{Type: graphql.String},
			"email":    &graphql.Field{Type: graphql.String},
			"salary":   &graphql.Field{Type: graphql.Float},
			"isActive": &graphql.Field{Type: graphql.Boolean},
			// Nested Resolver to link Employee to Department (Requirement #13)
			"department": &graphql.Field{
				Type: departmentType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					emp, ok := p.Source.(Employee)
					if !ok {
						return nil, fmt.Errorf("invalid employee source")
					}
					for _, d := range departments {
						if d.ID == emp.DepartmentID {
							return d, nil
						}
					}
					return nil, nil
				},
			},
		},
	})
}

// 10. Input Types for Mutations
var createEmployeeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateEmployeeInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"salary":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
		"departmentId": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
	},
})

var updateEmployeeInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateEmployeeInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":         &graphql.InputObjectFieldConfig{Type: graphql.String},
		"email":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"salary":       &graphql.InputObjectFieldConfig{Type: graphql.Float},
		"departmentId": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"isActive":     &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
	},
})

// ============================================================================
// ROOT QUERY & RESOLVERS (Requirements #4, #5, #11, #13, #14)
// ============================================================================

func createRootQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			// 5. Employee List Query
			"employees": &graphql.Field{
				Type: graphql.NewList(employeeType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return employees, nil
				},
			},
			// 4 & 11. Single Employee Query using Query Arguments
			"employee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				// 13. Resolvers & 14. Error Handling
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					for _, emp := range employees {
						if int(emp.ID) == id {
							return emp, nil
						}
					}
					// 14. Return custom GraphQL error
					return nil, fmt.Errorf("error: employee with id %d not found", id)
				},
			},
			"departments": &graphql.Field{
				Type: graphql.NewList(departmentType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return departments, nil
				},
			},
		},
	})
}

// ============================================================================
// ROOT MUTATION & RESOLVERS (Requirements #6, #7, #8, #13, #14)
// ============================================================================

func createRootMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			// 6. Create Employee Mutation (using Input Type)
			"createEmployee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(createEmployeeInputType)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})

					// Verify Department exists (14. Error Handling)
					deptID := input["departmentId"].(string)
					deptExists := false
					for _, d := range departments {
						if d.ID == deptID {
							deptExists = true
							break
						}
					}
					if !deptExists {
						return nil, fmt.Errorf("invalid departmentId: %s", deptID)
					}

					newEmp := Employee{
						ID:           int32(len(employees) + 1),
						Name:         input["name"].(string),
						Email:        input["email"].(string),
						Salary:       input["salary"].(float64),
						IsActive:     true,
						DepartmentID: deptID,
					}
					employees = append(employees, newEmp)
					return newEmp, nil
				},
			},
			// 7. Update Employee Mutation
			"updateEmployee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(updateEmployeeInputType)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					input := p.Args["input"].(map[string]interface{})

					for idx, emp := range employees {
						if int(emp.ID) == id {
							if val, ok := input["name"].(string); ok {
								employees[idx].Name = val
							}
							if val, ok := input["email"].(string); ok {
								employees[idx].Email = val
							}
							if val, ok := input["salary"].(float64); ok {
								employees[idx].Salary = val
							}
							if val, ok := input["departmentId"].(string); ok {
								employees[idx].DepartmentID = val
							}
							if val, ok := input["isActive"].(bool); ok {
								employees[idx].IsActive = val
							}
							return employees[idx], nil
						}
					}
					return nil, fmt.Errorf("employee with id %d not found", id)
				},
			},
			// 8. Delete Employee Mutation
			"deleteEmployee": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					for idx, emp := range employees {
						if int(emp.ID) == id {
							employees = append(employees, employees[:idx]...)
							employees = append(employees, employees[idx+1:]...)
							return true, nil
						}
					}
					return false, fmt.Errorf("employee with id %d not found", id)
				},
			},
		},
	})
}

// ============================================================================
// 1 & 15. CREATE SCHEMA AND SERVER
// ============================================================================

func createSchema() (graphql.Schema, error) {
	initTypes()
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    createRootQuery(),
		Mutation: createRootMutation(),
	})
}

func main() {
	schema, err := createSchema()
	if err != nil {
		panic(err)
	}

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", graphqlHandler)

	fmt.Println("Employee Management GraphQL Server Running")
	fmt.Println("GraphiQL UI: http://localhost:8080/graphql")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server stopped:", err)
	}
}