package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// =====================================================
// TASK 2: Create Employee type (Go struct)
// =====================================================
type Employee struct {
	ID           int
	Name         string
	Email        string
	Age          int
	Salary       float64
	DepartmentID int
}

// TASK 3: Create Department type (Go struct)
// =====================================================
type Department struct {
	ID   int
	Name string
}

// Sample in-memory data
var departments = []Department{
	{ID: 1, Name: "Engineering"},
	{ID: 2, Name: "HR"},
	{ID: 3, Name: "Finance"},
}

var employees = []Employee{
	{ID: 1, Name: "Ravi", Email: "ravi@example.com", Age: 25, Salary: 45000, DepartmentID: 1},
	{ID: 2, Name: "Neha", Email: "neha@example.com", Age: 27, Salary: 52000, DepartmentID: 2},
}

var nextEmployeeID = 3

// =====================================================
// Helper: Find employee by ID
// =====================================================
func findEmployee(id int) (Employee, bool) {
	for _, employee := range employees {
		if employee.ID == id {
			return employee, true
		}
	}
	return Employee{}, false
}

// =====================================================
// Helper: Find department by ID
// =====================================================
func findDepartment(id int) (Department, bool) {
	for _, department := range departments {
		if department.ID == id {
			return department, true
		}
	}
	return Department{}, false
}

// =====================================================
// TASK 14: Validate employee data and handle errors
// =====================================================
func validateEmployee(e Employee) error {
	if strings.TrimSpace(e.Name) == "" {
		return fmt.Errorf("employee name cannot be empty")
	}

	if !strings.Contains(e.Email, "@") {
		return fmt.Errorf("invalid email")
	}

	if e.Age <= 0 {
		return fmt.Errorf("age must be greater than zero")
	}

	if e.Salary < 0 {
		return fmt.Errorf("salary cannot be negative")
	}

	if _, found := findDepartment(e.DepartmentID); !found {
		return fmt.Errorf("department not found")
	}

	return nil
}

// Convert GraphQL ID into Go int
func parseID(value interface{}) (int, error) {
	switch v := value.(type) {
	case string:
		return strconv.Atoi(v)
	case int:
		return v, nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("invalid ID")
	}
}

// =====================================================
// TASK 10: Create helper for EmployeeInput
// =====================================================
func readEmployeeInput(value interface{}) (Employee, error) {
	input, ok := value.(map[string]interface{})
	if !ok {
		return Employee{}, fmt.Errorf("invalid employee input")
	}

	name, nameOK := input["name"].(string)
	email, emailOK := input["email"].(string)
	age, ageOK := input["age"].(int)
	salary, salaryOK := input["salary"].(float64)
	departmentID, departmentOK := input["departmentId"].(int)

	if !nameOK || !emailOK || !ageOK ||
		!salaryOK || !departmentOK {
		return Employee{}, fmt.Errorf("invalid input fields")
	}

	return Employee{
		Name:         name,
		Email:        email,
		Age:          age,
		Salary:       salary,
		DepartmentID: departmentID,
	}, nil
}

// =====================================================
// TASK 1: Create GraphQL schema
// TASK 2: Create Employee GraphQL type
// TASK 3: Create Department GraphQL type
// TASK 9: Use GraphQL scalar types
// TASK 10: Create input type
// TASK 11: Use query arguments
// TASK 12: Support GraphQL variables
// TASK 13: Implement resolvers
// TASK 14: GraphQL error handling
// =====================================================
func newGraphQLSchema() (graphql.Schema, error) {

	// TASK 3: Department GraphQL type
	departmentType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Department",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
	})

	// TASK 2: Employee GraphQL type
	// TASK 9: ID, String, Int and Float scalar types
	employeeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.NewNonNull(graphql.ID),
			},
			"name": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.Field{
				Type: graphql.NewNonNull(graphql.String),
			},
			"age": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"salary": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Float),
			},

			// Resolve employee's department
			"department": &graphql.Field{
				Type: departmentType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					employee := p.Source.(Employee)

					department, found := findDepartment(
						employee.DepartmentID,
					)

					if !found {
						return nil, nil
					}

					return department, nil
				},
			},
		},
	})

	// TASK 10: Employee input type
	employeeInputType := graphql.NewInputObject(
		graphql.InputObjectConfig{
			Name: "EmployeeInput",
			Fields: graphql.InputObjectConfigFieldMap{
				"name": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"email": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"age": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
				"salary": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.Float),
				},
				"departmentId": &graphql.InputObjectFieldConfig{
					Type: graphql.NewNonNull(graphql.Int),
				},
			},
		},
	)

	// =================================================
	// TASK 4: Implement single Employee query
	// TASK 5: Implement Employee list query
	// TASK 11: Use query arguments
	// =================================================
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{

			// Get one employee by ID
			"employee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := parseID(p.Args["id"])
					if err != nil {
						return nil, err
					}

					employee, found := findEmployee(id)
					if !found {
						return nil, fmt.Errorf(
							"employee with ID %d not found", id,
						)
					}

					return employee, nil
				},
			},

			// Get all employees
			"employees": &graphql.Field{
				Type: graphql.NewNonNull(
					graphql.NewList(
						graphql.NewNonNull(employeeType),
					),
				),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return employees, nil
				},
			},

			// Get all departments
			"departments": &graphql.Field{
				Type: graphql.NewNonNull(
					graphql.NewList(
						graphql.NewNonNull(departmentType),
					),
				),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return departments, nil
				},
			},
		},
	})

	// =================================================
	// TASK 6: Create Employee mutation
	// TASK 7: Update Employee mutation
	// TASK 8: Delete Employee mutation
	// =================================================
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{

			// TASK 6: Create employee
			"createEmployee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(employeeInputType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					employee, err := readEmployeeInput(
						p.Args["input"],
					)
					if err != nil {
						return nil, err
					}

					if err := validateEmployee(employee); err != nil {
						return nil, err
					}

					employee.ID = nextEmployeeID
					nextEmployeeID++

					employees = append(employees, employee)

					return employee, nil
				},
			},

			// TASK 7: Update employee
			"updateEmployee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
					"input": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(employeeInputType),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := parseID(p.Args["id"])
					if err != nil {
						return nil, err
					}

					updated, err := readEmployeeInput(
						p.Args["input"],
					)
					if err != nil {
						return nil, err
					}

					if err := validateEmployee(updated); err != nil {
						return nil, err
					}

					for i := range employees {
						if employees[i].ID == id {
							updated.ID = id
							employees[i] = updated
							return employees[i], nil
						}
					}

					return nil, fmt.Errorf(
						"employee with ID %d not found", id,
					)
				},
			},

			// TASK 8: Delete employee
			"deleteEmployee": &graphql.Field{
				Type: graphql.NewNonNull(graphql.Boolean),
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := parseID(p.Args["id"])
					if err != nil {
						return false, err
					}

					for i := range employees {
						if employees[i].ID == id {
							employees = append(
								employees[:i],
								employees[i+1:]...,
							)
							return true, nil
						}
					}

					return false, fmt.Errorf(
						"employee with ID %d not found", id,
					)
				},
			},
		},
	})

	// TASK 1: Combine queries and mutations into GraphQL schema
	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

// =====================================================
// TASK 15: Complete Employee GraphQL API
// Start HTTP server and expose GraphiQL interface
// =====================================================
func main() {
	schema, err := newGraphQLSchema()
	if err != nil {
		log.Fatal(err)
	}

	graphqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", graphqlHandler)

	fmt.Println("GraphQL server running at http://localhost:8080/graphql")

	log.Fatal(http.ListenAndServe(":8080", nil))
}


// {
//   employees{
//     id 
//     name
//     email
//     age
//     salary
//   }
// }
