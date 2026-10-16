
package resolver

import "employee-management/service"

type Resolver struct {
	employees   service.EmployeeService
	departments service.DepartmentService
	users       service.AuthService
}


func New(e service.EmployeeService, d service.DepartmentService, a service.AuthService) *Resolver {
	return &Resolver{employees: e, departments: d, users: a}
}
