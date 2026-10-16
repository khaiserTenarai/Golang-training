// Package resolver holds the code that runs when a GraphQL field is requested.
// The services are INJECTED (passed in) from main.go - resolvers never create them.
package resolver

import "employee-management/service"

type Resolver struct {
	employees   service.EmployeeService
	departments service.DepartmentService
	users       service.AuthService
}

// New receives the services it needs (dependency injection).
// In tests we pass fake services, in main.go the real ones.
func New(e service.EmployeeService, d service.DepartmentService, a service.AuthService) *Resolver {
	return &Resolver{employees: e, departments: d, users: a}
}
