package schema

import (
	"fmt"

	"employee-platform/models"
	"employee-platform/service"

	"github.com/graphql-go/graphql"
)

var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Department",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.String},
		"name":     &graphql.Field{Type: graphql.String},
		"location": &graphql.Field{Type: graphql.String},
	},
})

func BuildSchema(svc service.EmployeeService) (graphql.Schema, error) {
	EmployeeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"id":           &graphql.Field{Type: graphql.String},
			"name":         &graphql.Field{Type: graphql.String},
			"role":         &graphql.Field{Type: graphql.String},
			"departmentId": &graphql.Field{Type: graphql.String},
			"salary":       &graphql.Field{Type: graphql.Float},
			// Nested Employee -> Department resolver with optimized lookup
			"department": &graphql.Field{
				Type: DepartmentType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					emp, ok := p.Source.(models.Employee)
					if !ok {
						return nil, nil
					}
					return svc.GetDepartmentByID(emp.DepartmentID)
				},
			},
		},
	})

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employees": &graphql.Field{
				Type: graphql.NewList(EmployeeType),
				Args: graphql.FieldConfigArgument{
					"departmentId": &graphql.ArgumentConfig{Type: graphql.String},
					"minSalary":    &graphql.ArgumentConfig{Type: graphql.Float},
					"limit":        &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 10},
					"offset":       &graphql.ArgumentConfig{Type: graphql.Int, DefaultValue: 0},
					"sortBy":       &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "name"},
					"sortOrder":    &graphql.ArgumentConfig{Type: graphql.String, DefaultValue: "asc"},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					var filter models.EmployeeFilter
					if val, ok := p.Args["departmentId"].(string); ok {
						filter.DepartmentID = &val
					}
					if val, ok := p.Args["minSalary"].(float64); ok {
						filter.MinSalary = &val
					}

					page := models.Pagination{
						Limit:  p.Args["limit"].(int),
						Offset: p.Args["offset"].(int),
					}

					emps, _, err := svc.ListEmployees(filter, page, p.Args["sortBy"].(string), p.Args["sortOrder"].(string))
					return emps, err
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createEmployee": &graphql.Field{
				Type: EmployeeType,
				Args: graphql.FieldConfigArgument{
					"id":           &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"name":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"role":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"departmentId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"salary":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Authorization check inside GraphQL Resolver
					userCtx, ok := p.Context.Value("userCtx").(models.UserContext)
					if !ok {
						return nil, fmt.Errorf("unauthorized")
					}
					if userCtx.Role != "ADMIN" {
						return nil, fmt.Errorf("permission denied: ADMIN role required")
					}

					salary := p.Args["salary"].(float64)
					if salary <= 0 {
						return nil, fmt.Errorf("salary must be greater than 0")
					}

					emp := models.Employee{
						ID:           p.Args["id"].(string),
						Name:         p.Args["name"].(string),
						Role:         p.Args["role"].(string),
						DepartmentID: p.Args["departmentId"].(string),
						Salary:       salary,
					}

					if err := svc.CreateEmployee(&emp); err != nil {
						return nil, err
					}

					return emp, nil
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}
