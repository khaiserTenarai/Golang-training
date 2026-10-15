// Tasks 4-8, 11, 13, 14: Resolvers — the heart of GraphQL execution
// Task 4  — Employee query (single)
// Task 5  — Employee list query
// Task 6  — Create Employee mutation
// Task 7  — Update mutation
// Task 8  — Delete mutation
// Task 11 — Query arguments (id, limit, offset, etc.)
// Task 13 — Implement resolvers (all resolve functions)
// Task 14 — GraphQL error handling

package resolver

import (
	"fmt"

	"graphql_fundamentals/store"
	"graphql_fundamentals/types"

	"github.com/graphql-go/graphql"
)

// Task 13: Build resolvers with injected store
func BuildSchema(s *store.Store) (graphql.Schema, error) {

	// Add nested department field to Employee type
	types.EmployeeType.AddFieldConfig("department", &graphql.Field{
		Type:        types.DepartmentType,
		Description: "The employee's department (nested query)",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			emp, ok := p.Source.(store.Employee)
			if !ok {
				return nil, nil
			}
			dept, err := s.GetDepartment(emp.DepartmentID)
			if err != nil {
				return nil, nil // gracefully return null if dept not found
			}
			return dept, nil
		},
	})

	// ---- QUERIES ----
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			// Task 4: Single employee query
			// Task 11: Query argument — id (required)
			// Task 13: Resolver implementation
			// Task 14: Error handling (not found)
			"employee": &graphql.Field{
				Type:        types.EmployeeType,
				Description: "Get a single employee by ID",
				Args: graphql.FieldConfigArgument{
					// Task 11: Query argument
					"id": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(graphql.Int),
						Description: "Employee ID (required)",
					},
				},
				// Task 13: Resolver function
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if !ok {
						// Task 14: Error handling
						return nil, fmt.Errorf("invalid ID argument")
					}
					emp, err := s.GetEmployee(id)
					if err != nil {
						// Task 14: Error returned to client
						return nil, err
					}
					return emp, nil
				},
			},

			// Task 5: Employee list query
			// Task 11: Multiple query arguments (limit, offset, active filter)
			"employees": &graphql.Field{
				Type:        types.EmployeeListType,
				Description: "Get all employees with optional limit/offset",
				Args: graphql.FieldConfigArgument{
					// Task 11: Optional arguments with defaults
					"limit":     &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 10, Description: "Max results"},
					"offset":    &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 0, Description: "Skip N results"},
					"is_active": &graphql.ArgumentConfig{Type: graphql.Boolean, Description: "Filter by active status"},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					all := s.GetAllEmployees()
					limit := p.Args["limit"].(int)
					offset := p.Args["offset"].(int)

					// Filter by active status if provided
					var filtered []store.Employee
					if activeFilter, ok := p.Args["is_active"].(bool); ok {
						for _, e := range all {
							if e.IsActive == activeFilter {
								filtered = append(filtered, e)
							}
						}
					} else {
						filtered = all
					}

					// Apply offset/limit
					total := len(filtered)
					if offset > total {
						offset = total
					}
					end := offset + limit
					if end > total {
						end = total
					}

					return map[string]interface{}{
						"employees": filtered[offset:end],
						"count":     total,
					}, nil
				},
			},

			// Department query
			"department": &graphql.Field{
				Type: types.DepartmentType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					return s.GetDepartment(id)
				},
			},

			// Department list
			"departments": &graphql.Field{
				Type:        graphql.NewList(types.DepartmentType),
				Description: "Get all departments",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return s.GetAllDepartments(), nil
				},
			},
		},
	})

	// ---- MUTATIONS ----
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{

			// Task 6: Create Employee mutation
			// Task 10: Uses CreateEmployeeInput type
			// Task 14: Validation error handling
			"createEmployee": &graphql.Field{
				Type:        types.EmployeeType,
				Description: "Create a new employee",
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type:        graphql.NewNonNull(types.CreateEmployeeInput),
						Description: "Employee data (Task 10: input type)",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					emp := store.Employee{
						Name:         input["name"].(string),
						Email:        input["email"].(string),
						Age:          input["age"].(int),
						DepartmentID: input["department_id"].(int),
						Salary:       input["salary"].(float64),
						IsActive:     true,
					}
					if v, ok := input["is_active"].(bool); ok {
						emp.IsActive = v
					}
					if v, ok := input["join_date"].(string); ok && v != "" {
						emp.JoinDate = v
					}
					// Task 14: Validation errors returned
					created, err := s.CreateEmployee(emp)
					if err != nil {
						return nil, err
					}
					return created, nil
				},
			},

			// Task 7: Update mutation
			"updateEmployee": &graphql.Field{
				Type:        types.EmployeeType,
				Description: "Update an existing employee",
				Args: graphql.FieldConfigArgument{
					"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(types.UpdateEmployeeInput)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					input := p.Args["input"].(map[string]interface{})
					emp := store.Employee{
						Name:         input["name"].(string),
						Email:        input["email"].(string),
						Age:          input["age"].(int),
						DepartmentID: input["department_id"].(int),
						Salary:       input["salary"].(float64),
						IsActive:     true,
					}
					if v, ok := input["is_active"].(bool); ok {
						emp.IsActive = v
					}
					if v, ok := input["join_date"].(string); ok {
						emp.JoinDate = v
					}
					return s.UpdateEmployee(id, emp)
				},
			},

			// Task 8: Delete mutation
			"deleteEmployee": &graphql.Field{
				Type:        types.DeleteResultType,
				Description: "Delete an employee by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(int)
					err := s.DeleteEmployee(id)
					if err != nil {
						return map[string]interface{}{
							"success": false,
							"message": err.Error(),
						}, nil // Return error as data, not GraphQL error
					}
					return map[string]interface{}{
						"success": true,
						"message": fmt.Sprintf("Employee %d deleted", id),
					}, nil
				},
			},
		},
	})

	// Task 1: Create the GraphQL schema
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
