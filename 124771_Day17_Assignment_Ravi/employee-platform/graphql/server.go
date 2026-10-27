// Package graphql contains the GraphQL schema and resolvers.
// Tasks 1-5 and 8-11 are implemented here and in the shared service/repository layers.
package graphql

import (
	"context"
	"fmt"
	"strings"

	"employee-platform/model"
	"employee-platform/service"
	"employee-platform/utils"

	gql "github.com/graphql-go/graphql"
)

type claimsContextKey struct{}

type Server struct {
	// Task 3: services are injected instead of being constructed inside resolvers.
	employeeService   *service.EmployeeService
	departmentService *service.DepartmentService
	authService       *service.AuthService
	schema            gql.Schema
}

func NewServer(employeeService *service.EmployeeService, departmentService *service.DepartmentService, authService *service.AuthService) (*Server, error) {
	s := &Server{employeeService: employeeService, departmentService: departmentService, authService: authService}
	schema, err := s.buildSchema()
	if err != nil {
		return nil, err
	}
	s.schema = schema
	return s, nil
}

func (s *Server) buildSchema() (gql.Schema, error) {
	// Task 1: schema objects are defined separately and can be moved into schema modules.
	departmentType := gql.NewObject(gql.ObjectConfig{Name: "Department", Fields: gql.Fields{
		"id":          &gql.Field{Type: gql.NewNonNull(gql.Int)},
		"name":        &gql.Field{Type: gql.NewNonNull(gql.String)},
		"description": &gql.Field{Type: gql.String},
	}})

	// Task 4: Employee has a nested department field.
	employeeType := gql.NewObject(gql.ObjectConfig{Name: "Employee", Fields: gql.Fields{
		"id":     &gql.Field{Type: gql.NewNonNull(gql.Int)},
		"name":   &gql.Field{Type: gql.NewNonNull(gql.String)},
		"email":  &gql.Field{Type: gql.NewNonNull(gql.String)},
		"salary": &gql.Field{Type: gql.NewNonNull(gql.Float)},
		"departmentId": &gql.Field{Type: gql.NewNonNull(gql.Int), Resolve: func(p gql.ResolveParams) (interface{}, error) {
			e, ok := asEmployee(p.Source)
			if !ok {
				return 0, nil
			}
			return int(e.DepartmentID), nil
		}},
		"department": &gql.Field{Type: departmentType, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			e, ok := asEmployee(p.Source)
			if !ok {
				return nil, nil
			}
			return s.departmentService.ByID(e.DepartmentID)
		}},
	}})

	// Task 8: GraphQL non-null types validate required input fields; service layer validates business rules.
	employeeInput := gql.NewInputObject(gql.InputObjectConfig{Name: "EmployeeInput", Fields: gql.InputObjectConfigFieldMap{
		"name":         &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"email":        &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"salary":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		"departmentId": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Int)},
	}})

	pageType := gql.NewObject(gql.ObjectConfig{Name: "EmployeePage", Fields: gql.Fields{
		"items": &gql.Field{Type: gql.NewList(employeeType)},
		"total": &gql.Field{Type: gql.Int}, "page": &gql.Field{Type: gql.Int}, "limit": &gql.Field{Type: gql.Int},
	}})

	// Task 2: resolver functions are grouped in this GraphQL package.
	query := gql.NewObject(gql.ObjectConfig{Name: "Query", Fields: gql.Fields{
		"employee": &gql.Field{Type: employeeType, Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if _, err := requireClaims(p.Context); err != nil {
				return nil, err
			}
			id := uint(intArg(p.Args, "id", 0))
			if id == 0 {
				return nil, fmt.Errorf("id must be positive")
			}
			return s.employeeService.ByID(id)
		}},
		"employees": &gql.Field{Type: pageType, Args: gql.FieldConfigArgument{
			"page": &gql.ArgumentConfig{Type: gql.Int}, "limit": &gql.ArgumentConfig{Type: gql.Int},
			"departmentId": &gql.ArgumentConfig{Type: gql.Int}, "search": &gql.ArgumentConfig{Type: gql.String},
			"sortBy": &gql.ArgumentConfig{Type: gql.String}, "order": &gql.ArgumentConfig{Type: gql.String},
		}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if _, err := requireClaims(p.Context); err != nil {
				return nil, err
			}
			return s.employeeService.List(intArg(p.Args, "page", 1), intArg(p.Args, "limit", 10), uint(intArg(p.Args, "departmentId", 0)),
				stringArg(p.Args, "search"), stringArgDefault(p.Args, "sortBy", "id"), stringArgDefault(p.Args, "order", "asc"))
		}},
		"departments": &gql.Field{Type: gql.NewList(departmentType), Resolve: func(p gql.ResolveParams) (interface{}, error) {
			if _, err := requireClaims(p.Context); err != nil {
				return nil, err
			}
			return s.departmentService.List()
		}},
	}})

	mutation := gql.NewObject(gql.ObjectConfig{Name: "Mutation", Fields: gql.Fields{
		"createEmployee": &gql.Field{Type: employeeType, Args: gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)}}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			claims, err := requireClaims(p.Context)
			if err != nil {
				return nil, err
			}
			// Task 7: only admins may create employees.
			if claims.Role != "admin" {
				return nil, fmt.Errorf("permission denied: admin role required")
			}
			in := p.Args["input"].(map[string]interface{})
			return s.employeeService.Create(strings.TrimSpace(in["name"].(string)), strings.TrimSpace(in["email"].(string)), in["salary"].(float64), uint(in["departmentId"].(int)))
		}},
		"updateEmployee": &gql.Field{Type: employeeType, Args: gql.FieldConfigArgument{
			"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}, "input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)},
		}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			claims, err := requireClaims(p.Context)
			if err != nil {
				return nil, err
			}
			if claims.Role != "admin" {
				return nil, fmt.Errorf("permission denied: admin role required")
			}
			in := p.Args["input"].(map[string]interface{})
			return s.employeeService.Update(uint(intArg(p.Args, "id", 0)), in["name"].(string), in["email"].(string), in["salary"].(float64), uint(in["departmentId"].(int)))
		}},
		"deleteEmployee": &gql.Field{Type: gql.NewNonNull(gql.Boolean), Args: gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.Int)}}, Resolve: func(p gql.ResolveParams) (interface{}, error) {
			claims, err := requireClaims(p.Context)
			if err != nil {
				return false, err
			}
			if claims.Role != "admin" {
				return false, fmt.Errorf("permission denied: admin role required")
			}
			id := uint(intArg(p.Args, "id", 0))
			if id == 0 {
				return false, fmt.Errorf("id must be positive")
			}
			if err := s.employeeService.Delete(id); err != nil {
				return false, err
			}
			return true, nil
		}},
	}})

	return gql.NewSchema(gql.SchemaConfig{Query: query, Mutation: mutation})
}

func asEmployee(source interface{}) (model.Employee, bool) {
	switch e := source.(type) {
	case model.Employee:
		return e, true
	case *model.Employee:
		if e != nil {
			return *e, true
		}
	}
	return model.Employee{}, false
}
func intArg(args map[string]interface{}, key string, fallback int) int {
	if v, ok := args[key].(int); ok {
		return v
	}
	return fallback
}
func stringArg(args map[string]interface{}, key string) string { v, _ := args[key].(string); return v }
func stringArgDefault(args map[string]interface{}, key, fallback string) string {
	if v := stringArg(args, key); v != "" {
		return v
	}
	return fallback
}
func requireClaims(ctx context.Context) (*utils.Claims, error) {
	claims, ok := ctx.Value(claimsContextKey{}).(*utils.Claims)
	if !ok || claims == nil {
		return nil, fmt.Errorf("authentication required: provide Authorization: Bearer <token>")
	}
	return claims, nil
}
