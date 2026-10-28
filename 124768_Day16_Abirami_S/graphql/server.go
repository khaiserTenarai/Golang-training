package graphql

import (
	"errors"
	"fmt"
	"strings"

	"day16-employee-graphql/model"
	"day16-employee-graphql/repository"
	"day16-employee-graphql/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	service service.EmployeeService
	schema  gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	x := &Server{service: s}
	sc, err := x.buildSchema()
	if err != nil {
		return nil, err
	}
	x.schema = sc
	return x, nil
}

func (x *Server) buildSchema() (gql.Schema, error) {
	department := gql.NewObject(gql.ObjectConfig{
		Name: "Department",
		Fields: gql.Fields{
			"id":   &gql.Field{Type: gql.NewNonNull(gql.ID)},
			"name": &gql.Field{Type: gql.NewNonNull(gql.String)},
		},
	})

	emp := gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":     &gql.Field{Type: gql.NewNonNull(gql.ID)},
			"name":   &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email":  &gql.Field{Type: gql.NewNonNull(gql.String)},
			"age":    &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"salary": &gql.Field{Type: gql.NewNonNull(gql.Float)},
			"department": &gql.Field{
				Type: gql.NewNonNull(department),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					e, ok := p.Source.(model.Employee)
					if !ok {
						if ep, ok := p.Source.(*model.Employee); ok {
							e = *ep
						} else {
							return nil, fmt.Errorf("invalid employee source")
						}
					}

					departments := map[int]model.Department{
						1: {ID: 1, Name: "Engineering"},
						2: {ID: 2, Name: "HR"},
						3: {ID: 3, Name: "Finance"},
					}

					d, ok := departments[e.DepartmentID]
					if !ok {
						return nil, fmt.Errorf("department not found")
					}
					return d, nil
				},
			},
		},
	})

	input := gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":         &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email":        &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"age":          &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Int)},
			"departmentId": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.ID)},
			"salary":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		},
	})

	query := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employee": &gql.Field{
				Type: emp,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseID(p.Args["id"])
					if err != nil {
						return nil, err
					}
					e, err := x.service.GetEmployee(id)
					if err != nil {
						if errors.Is(err, repository.ErrEmployeeNotFound) {
							return nil, fmt.Errorf("employee not found")
						}
						return nil, err
					}
					return e, nil
				},
			},
			"employees": &gql.Field{
				Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(emp))),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return x.service.GetEmployees()
				},
			},
		},
	})

	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: gql.NewNonNull(emp),
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(input)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					e := fromInput(p.Args["input"].(map[string]interface{}))
					return x.service.CreateEmployee(e)
				},
			},
			"updateEmployee": &gql.Field{
				Type: gql.NewNonNull(emp),
				Args: gql.FieldConfigArgument{
					"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
					"input": &gql.ArgumentConfig{Type: gql.NewNonNull(input)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseID(p.Args["id"])
					if err != nil {
						return nil, err
					}
					e := fromInput(p.Args["input"].(map[string]interface{}))
					return x.service.UpdateEmployee(id, e)
				},
			},
			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseID(p.Args["id"])
					if err != nil {
						return false, err
					}
					if _, err := x.service.DeleteEmployee(id); err != nil {
						return false, err
					}
					return true, nil
				},
			},
		},
	})

	return gql.NewSchema(gql.SchemaConfig{
		Query:    query,
		Mutation: mutation,
	})
}

func parseID(value interface{}) (int, error) {
	switch v := value.(type) {
	case int:
		return v, nil
	case string:
		var id int
		_, err := fmt.Sscanf(v, "%d", &id)
		return id, err
	default:
		return 0, fmt.Errorf("invalid employee id")
	}
}

func fromInput(m map[string]interface{}) model.Employee {
	return model.Employee{
		Name:         strings.TrimSpace(fmt.Sprint(m["name"])),
		Email:        strings.TrimSpace(fmt.Sprint(m["email"])),
		Age:          m["age"].(int),
		DepartmentID: mustParseID(m["departmentId"]),
		Salary:       m["salary"].(float64),
	}
}

func mustParseID(value interface{}) int {
	id, _ := parseID(value)
	return id
}
