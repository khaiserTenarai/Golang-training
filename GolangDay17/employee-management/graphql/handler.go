package graphql

import (
	"employee-management/graphql/resolver"
	"employee-management/graphql/schema"
	"employee-management/service"
	"github.com/graphql-go/handler"
	"net/http"
)


func NewHTTPHandler(e service.EmployeeService, d service.DepartmentService, a service.AuthService) (http.Handler, error) {
	r := resolver.New(e, d, a) 
	s, err := schema.Build(r)
	if err != nil {
		return nil, err
	}
	return handler.New(&handler.Config{Schema: &s, Pretty: true, GraphiQL: false, Playground: true}), nil
}