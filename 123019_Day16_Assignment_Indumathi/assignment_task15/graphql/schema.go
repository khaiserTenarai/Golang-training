package graphql

import (
	"fmt"

	"employee-management/model"
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	service service.EmployeeService
	schema  gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	srv := &Server{service: s}
	sc, err := srv.buildSchema()
	if err != nil {
		return nil, err
	}
	srv.schema = sc
	return srv, nil
}

func (s *Server) buildSchema() (gql.Schema, error) {
	// 3. Department Type
	departmentType := gql.NewObject(gql.ObjectConfig{
		Name: "Department",
		Fields: gql.Fields{
			"id":   &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"name": &gql.Field{Type: gql.NewNonNull(gql.String)},
		},
	})

	// 2. Employee Type (uses Scalar Types: Int, String, Float)
	employeeType := gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":         &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"name":       &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.Field{Type: gql.NewNonNull(gql.String)},
			"department": &gql.Field{Type: gql.NewNonNull(departmentType)},
			"salary":     &gql.Field{Type: gql.NewNonNull(gql.Float)},
		},
	})

	// 10. Input Types
	departmentInput := gql.NewInputObject(gql.InputObjectConfig{
		Name: "DepartmentInput",
		Fields: gql.InputObjectConfigFieldMap{
			"id":   &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Int)},
			"name": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		},
	})

	employeeInput := gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(departmentInput)},
			"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		},
	})

	// 4, 5, 11. Queries (Employee & Employee List with Query Arguments)
	queryType := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
				},
				// 13, 14. Resolvers & Error Handling
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := int64(p.Args["id"].(int))
					return s.service.GetByID(p.Context, id)
				},
			},
			"employees": &gql.Field{
				Type: gql.NewList(employeeType),
				Args: gql.FieldConfigArgument{
					"sort": &gql.ArgumentConfig{Type: gql.String},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					sortBy, _ := p.Args["sort"].(string)
					return s.service.GetAll(p.Context, sortBy)
				},
			},
		},
	})

	// 6, 7, 8. Mutations (Create, Update, Delete)
	mutationType := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					emp := parseEmployeeInput(p.Args["input"].(map[string]interface{}))
					if err := s.service.Create(p.Context, &emp); err != nil {
						return nil, err
					}
					return emp, nil
				},
			},
			"updateEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := int64(p.Args["id"].(int))
					emp := parseEmployeeInput(p.Args["input"].(map[string]interface{}))
					return s.service.Update(p.Context, id, &emp)
				},
			},
			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := int64(p.Args["id"].(int))
					if err := s.service.Delete(p.Context, id); err != nil {
						return false, err
					}
					return true, nil
				},
			},
		},
	})

	// 1. Create GraphQL Schema
	return gql.NewSchema(gql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

func parseEmployeeInput(m map[string]interface{}) model.Employee {
	deptMap := m["department"].(map[string]interface{})
	return model.Employee{
		Name:  fmt.Sprint(m["name"]),
		Email: fmt.Sprint(m["email"]),
		Department: model.Department{
			ID:   deptMap["id"].(int),
			Name: fmt.Sprint(deptMap["name"]),
		},
		Salary: m["salary"].(float64),
	}
}