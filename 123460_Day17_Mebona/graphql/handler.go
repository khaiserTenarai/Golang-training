package graphql

import (
	"employee-management/graphql/resolver"
	"employee-management/graphql/schema"
	"employee-management/service"
	"github.com/graphql-go/handler"
	"net/http"
)

// NewHTTPHandler wires everything together:
// services -> resolver -> schema -> HTTP handler (with the Playground UI).
func NewHTTPHandler(e service.EmployeeService, d service.DepartmentService, a service.AuthService) (http.Handler, error) {
	r := resolver.New(e, d, a) // dependency injection
	s, err := schema.Build(r)
	if err != nil {
		return nil, err
	}
	return handler.New(&handler.Config{Schema: &s, Pretty: true, GraphiQL: false, Playground: true}), nil
}
