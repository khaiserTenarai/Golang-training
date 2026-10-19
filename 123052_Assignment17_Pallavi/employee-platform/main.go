package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// ============================================================================
// MODELS
// ============================================================================
type Department struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	DepartmentID int    `json:"department_id"`
	Salary       int    `json:"salary"`
}

type UserContext struct {
	Role string
}

// In-Memory Storage
var departments = []Department{
	{ID: 101, Name: "Engineering"},
	{ID: 102, Name: "Human Resources"},
}

var employees = []Employee{
	{ID: 1, Name: "Rajesh", Email: "rajesh@example.com", DepartmentID: 101, Salary: 60000},
	{ID: 2, Name: "Amit", Email: "amit@example.com", DepartmentID: 102, Salary: 50000},
	{ID: 3, Name: "Priya", Email: "priya@example.com", DepartmentID: 101, Salary: 75000},
}

// ============================================================================
// SERVICE LAYER (Business Logic, Filtering, Sorting, Pagination)
// ============================================================================
func GetDepartments() []Department {
	return departments
}

func GetDepartmentByID(id int) (*Department, error) {
	for _, dept := range departments {
		if dept.ID == id {
			return &dept, nil
		}
	}
	return nil, errors.New("department not found")
}

func GetEmployees(deptFilter string, sortBy string, limit int, offset int) []Employee {
	var result []Employee

	// 1. Filtering
	for _, emp := range employees {
		dept, _ := GetDepartmentByID(emp.DepartmentID)
		if deptFilter == "" || (dept != nil && strings.EqualFold(dept.Name, deptFilter)) {
			result = append(result, emp)
		}
	}

	// 2. Sorting
	if sortBy == "salary_asc" {
		sort.Slice(result, func(i, j int) bool { return result[i].Salary < result[j].Salary })
	} else if sortBy == "salary_desc" {
		sort.Slice(result, func(i, j int) bool { return result[i].Salary > result[j].Salary })
	}

	// 3. Pagination
	if offset >= len(result) {
		return []Employee{}
	}
	end := offset + limit
	if limit <= 0 || end > len(result) {
		end = len(result)
	}

	return result[offset:end]
}

func CreateEmployee(name, email string, deptID int, salary int) (*Employee, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("name is required")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("valid email is required")
	}
	if salary < 0 {
		return nil, errors.New("salary cannot be negative")
	}
	if _, err := GetDepartmentByID(deptID); err != nil {
		return nil, errors.New("invalid department_id")
	}

	newEmp := Employee{
		ID:           len(employees) + 1,
		Name:         name,
		Email:        email,
		DepartmentID: deptID,
		Salary:       salary,
	}
	employees = append(employees, newEmp)
	return &newEmp, nil
}

// ============================================================================
// MIDDLEWARE (Authentication & Context)
// ============================================================================
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		role := "user"
		if token == "Bearer admin-token" {
			role = "admin"
		}
		ctx := context.WithValue(r.Context(), "user", UserContext{Role: role})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ============================================================================
// GRAPHQL SCHEMA & RESOLVERS
// ============================================================================
var departmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":   &graphql.Field{Type: graphql.Int},
		"name": &graphql.Field{Type: graphql.String},
	},
})

var employeeType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Employee",
	Fields: graphql.Fields{
		"id":     &graphql.Field{Type: graphql.Int},
		"name":   &graphql.Field{Type: graphql.String},
		"email":  &graphql.Field{Type: graphql.String},
		"salary": &graphql.Field{Type: graphql.Int},
		// Nested Query: Employee -> Department
		"department": &graphql.Field{
			Type: departmentType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				emp := p.Source.(Employee)
				return GetDepartmentByID(emp.DepartmentID)
			},
		},
	},
})

func BuildGraphQLSchema() (graphql.Schema, error) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employees": &graphql.Field{
				Type: graphql.NewList(employeeType),
				Args: graphql.FieldConfigArgument{
					"department": &graphql.ArgumentConfig{Type: graphql.String},
					"sortBy":     &graphql.ArgumentConfig{Type: graphql.String},
					"limit":      &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 10},
					"offset":     &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 0},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					dept, _ := p.Args["department"].(string)
					sortBy, _ := p.Args["sortBy"].(string)
					limit, _ := p.Args["limit"].(int)
					offset, _ := p.Args["offset"].(int)

					return GetEmployees(dept, sortBy, limit, offset), nil
				},
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createEmployee": &graphql.Field{
				Type: employeeType,
				Args: graphql.FieldConfigArgument{
					"name":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"email":        &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"departmentId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
					"salary":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Authorization Check
					u, ok := p.Context.Value("user").(UserContext)
					if !ok || u.Role != "admin" {
						return nil, errors.New("unauthorized: admin privileges required")
					}

					return CreateEmployee(
						p.Args["name"].(string),
						p.Args["email"].(string),
						p.Args["departmentId"].(int),
						p.Args["salary"].(int),
					)
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}

// ============================================================================
// REST CONTROLLER
// ============================================================================
func RESTGetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dept := r.URL.Query().Get("department")
	sort := r.URL.Query().Get("sortBy")

	json.NewEncoder(w).Encode(GetEmployees(dept, sort, 10, 0))
}

// ============================================================================
// MAIN
// ============================================================================
func main() {
	schema, err := BuildGraphQLSchema()
	if err != nil {
		panic(err)
	}

	gqlHandler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.HandleFunc("/api/employees", RESTGetEmployees)
	http.Handle("/graphql", AuthMiddleware(gqlHandler))

	fmt.Println("Server running on http://localhost:8080")
	fmt.Println("GraphQL UI: http://localhost:8080/graphql")
	_ = http.ListenAndServe(":8080", nil)
}