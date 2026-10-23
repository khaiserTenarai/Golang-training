package graphql

import (
	"employee-management/model"
	"employee-management/service"
	"fmt"
	"strconv"
	"strings"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	service service.EmployeeService
	schema  gql.Schema
}

func NewServer(s service.EmployeeService) (*Server, error) {
	x := &Server{service: s}
	sc, err := x.buildSchema()
	if err != nil { return nil, err }
	x.schema = sc
	return x, nil
}

func (x *Server) buildSchema() (gql.Schema, error) {
	// Department is a GraphQL object type. The current database model stores
	// the department name as a string, so the resolver wraps it as { name: ... }.
	departmentType := gql.NewObject(gql.ObjectConfig{
		Name: "Department",
		Fields: gql.Fields{
			"name": &gql.Field{Type: gql.NewNonNull(gql.String)},
		},
	})

	employeeType := gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id": &gql.Field{Type: gql.NewNonNull(gql.ID)},
			"name": &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email": &gql.Field{Type: gql.NewNonNull(gql.String)},
			"department": &gql.Field{Type: gql.NewNonNull(departmentType), Resolve: func(p gql.ResolveParams) (interface{}, error) {
				e, ok := p.Source.(model.Employee)
				if !ok { if ptr, yes := p.Source.(*model.Employee); yes && ptr != nil { e = *ptr } else { return nil, fmt.Errorf("could not resolve employee department") } }
				return map[string]interface{}{"name": e.Department}, nil
			}},
			"salary": &gql.Field{Type: gql.NewNonNull(gql.Float)},
		},
	})

	employeeInput := gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"salary": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		},
	})

	query := gql.NewObject(gql.ObjectConfig{Name: "Query", Fields: gql.Fields{
		"employee": &gql.Field{
			Type: employeeType,
			Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)}},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				id, err := parseID(p.Args["id"]); if err != nil { return nil, err }
				e, err := x.service.GetByID(id)
				if service.IsNotFound(err) { return nil, fmt.Errorf("employee with id %d not found", id) }
				if err != nil { return nil, fmt.Errorf("get employee: %w", err) }
				return e, nil
			},
		},
		"employees": &gql.Field{Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(employeeType))), Resolve: func(p gql.ResolveParams) (interface{}, error) {
			employees, err := x.service.GetAll()
			if err != nil { return nil, fmt.Errorf("list employees: %w", err) }
			return employees, nil
		}},
	}})

	mutation := gql.NewObject(gql.ObjectConfig{Name: "Mutation", Fields: gql.Fields{
		"createEmployee": &gql.Field{
			Type: employeeType,
			Args: gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)}},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				m, ok := p.Args["input"].(map[string]interface{}); if !ok { return nil, fmt.Errorf("invalid employee input") }
				e, err := fromInput(m); if err != nil { return nil, err }
				if err := x.service.Create(&e); err != nil { return nil, fmt.Errorf("create employee: %w", err) }
				return e, nil
			},
		},
		"updateEmployee": &gql.Field{
			Type: employeeType,
			Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)}, "input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)}},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				id, err := parseID(p.Args["id"]); if err != nil { return nil, err }
				m, ok := p.Args["input"].(map[string]interface{}); if !ok { return nil, fmt.Errorf("invalid employee input") }
				e, err := fromInput(m); if err != nil { return nil, err }
				updated, err := x.service.Update(id, &e)
				if service.IsNotFound(err) { return nil, fmt.Errorf("employee with id %d not found", id) }
				if err != nil { return nil, fmt.Errorf("update employee: %w", err) }
				return updated, nil
			},
		},
		"deleteEmployee": &gql.Field{
			Type: gql.NewNonNull(gql.Boolean),
			Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)}},
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				id, err := parseID(p.Args["id"]); if err != nil { return false, err }
				if err := x.service.Delete(id); err != nil {
					if service.IsNotFound(err) { return false, fmt.Errorf("employee with id %d not found", id) }
					return false, fmt.Errorf("delete employee: %w", err)
				}
				return true, nil
			},
		},
	}})

	return gql.NewSchema(gql.SchemaConfig{Query: query, Mutation: mutation})
}

func parseID(value interface{}) (uint, error) {
	s := strings.TrimSpace(fmt.Sprint(value))
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil || n == 0 { return 0, fmt.Errorf("id must be a positive integer") }
	return uint(n), nil
}

func fromInput(m map[string]interface{}) (model.Employee, error) {
	name := strings.TrimSpace(fmt.Sprint(m["name"]))
	email := strings.TrimSpace(fmt.Sprint(m["email"]))
	department := strings.TrimSpace(fmt.Sprint(m["department"]))
	salary, ok := m["salary"].(float64)
	if !ok { return model.Employee{}, fmt.Errorf("salary must be a number") }
	if name == "" { return model.Employee{}, fmt.Errorf("name is required") }
	if email == "" || !strings.Contains(email, "@") { return model.Employee{}, fmt.Errorf("valid email is required") }
	if department == "" { return model.Employee{}, fmt.Errorf("department is required") }
	if salary < 0 { return model.Employee{}, fmt.Errorf("salary cannot be negative") }
	return model.Employee{Name: name, Email: email, Department: department, Salary: salary}, nil
}
