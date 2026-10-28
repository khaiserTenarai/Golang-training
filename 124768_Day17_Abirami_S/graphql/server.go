package graphql

import (
	"context"
	"net/http"
	"strings"

	"day17-employee-platform/middleware"
	"day17-employee-platform/model"
	"day17-employee-platform/service"

	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type Server struct {
	EmployeeService   service.EmployeeService
	DepartmentService service.DepartmentService
	Schema            gql.Schema
}

type contextKey string

const roleContextKey contextKey = "role"

func NewServer(es service.EmployeeService, ds service.DepartmentService) (*Server, error) {
	s := &Server{
		EmployeeService:   es,
		DepartmentService: ds,
	}

	department := departmentType()
	employee := employeeType(s, department)

	schema, err := gql.NewSchema(gql.SchemaConfig{
		Query:    buildQuery(s, employee, department),
		Mutation: buildMutation(s, employee),
	})
	if err != nil {
		return nil, err
	}

	s.Schema = schema
	return s, nil
}

func (s *Server) HTTPHandler() http.Handler {
	return handler.New(&handler.Config{
		Schema:   &s.Schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func WithOptionalRole(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		if tokenString == "" {
			next.ServeHTTP(w, r)
			return
		}

		role, ok := middleware.ParseRole(tokenString, secret)
		if !ok {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), roleContextKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WithRole(next http.Handler, secret string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))
		role, ok := middleware.ParseRole(tokenString, secret)
		if !ok {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), roleContextKey, role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func WithRoleContext(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, roleContextKey, role)
}

func roleFromContext(ctx context.Context) string {
	role, _ := ctx.Value(roleContextKey).(string)
	return role
}

func inputToEmployee(input map[string]interface{}) *model.Employee {
	e := &model.Employee{}

	if value, ok := input["name"].(string); ok {
		e.Name = value
	}
	if value, ok := input["email"].(string); ok {
		e.Email = value
	}
	if value, ok := input["salary"].(float64); ok {
		e.Salary = value
	}
	if value, ok := input["departmentId"].(int); ok {
		e.DepartmentID = uint(value)
	}

	return e
}
