package graphql

import (
	"employee-management/model"
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	service *service.EmployeeService
	schema  gql.Schema
}

func NewServer(svc *service.EmployeeService) (*Server, error) {
	server := &Server{service: svc}

	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query:    server.buildQuery(),
		Mutation: server.buildMutation(),
	})
	if err != nil {
		return nil, err
	}

	server.schema = schema
	return server, nil
}

var employeeType = gql.NewObject(gql.ObjectConfig{
	Name: "Employee",
	Fields: gql.Fields{
		"id":         &gql.Field{Type: gql.Int},
		"name":       &gql.Field{Type: gql.String},
		"email":      &gql.Field{Type: gql.String},
		"department": &gql.Field{Type: gql.String},
		"salary":     &gql.Field{Type: gql.Float},
	},
})

var employeeInput = gql.NewInputObject(gql.InputObjectConfig{
	Name: "EmployeeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
	},
})

func (s *Server) buildQuery() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			// Questions 9, 10 and 11
			"employees": &gql.Field{
				Type: gql.NewList(employeeType),
				Args: gql.FieldConfigArgument{
					"page": &gql.ArgumentConfig{
						Type:         gql.Int,
						DefaultValue: 1,
					},
					"limit": &gql.ArgumentConfig{
						Type:         gql.Int,
						DefaultValue: 10,
					},
					"department": &gql.ArgumentConfig{
						Type: gql.String,
					},
					"sortBy": &gql.ArgumentConfig{
						Type:         gql.String,
						DefaultValue: "id",
					},
					"sortOrder": &gql.ArgumentConfig{
						Type:         gql.String,
						DefaultValue: "asc",
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					page := intArg(p.Args, "page", 1)
					limit := intArg(p.Args, "limit", 10)
					department := stringArg(p.Args, "department", "")
					sortBy := stringArg(p.Args, "sortBy", "id")
					sortOrder := stringArg(p.Args, "sortOrder", "asc")

					return s.service.GetEmployees(
						page,
						limit,
						department,
						sortBy,
						sortOrder,
					)
				},
			},

			"employee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := uint(p.Args["id"].(int))
					return s.service.GetEmployee(id)
				},
			},
		},
	})
}

func (s *Server) buildMutation() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(employeeInput),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})

					employee := model.Employee{
						Name:       input["name"].(string),
						Email:      input["email"].(string),
						Department: input["department"].(string),
						Salary:     input["salary"].(float64),
					}

					return s.service.CreateEmployee(employee)
				},
			},

			"updateEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(employeeInput),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := uint(p.Args["id"].(int))
					input := p.Args["input"].(map[string]interface{})

					employee := model.Employee{
						Name:       input["name"].(string),
						Email:      input["email"].(string),
						Department: input["department"].(string),
						Salary:     input["salary"].(float64),
					}

					return s.service.UpdateEmployee(id, employee)
				},
			},

			"deleteEmployee": &gql.Field{
				Type: gql.Boolean,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.Int),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := uint(p.Args["id"].(int))
					return s.service.DeleteEmployee(id)
				},
			},
		},
	})
}

func intArg(args map[string]interface{}, key string, fallback int) int {
	value, ok := args[key].(int)
	if !ok {
		return fallback
	}
	return value
}

func stringArg(args map[string]interface{}, key string, fallback string) string {
	value, ok := args[key].(string)
	if !ok {
		return fallback
	}
	return value
}
