package resolver

import (
	"employee-management/model"
	gql "github.com/graphql-go/graphql"
)

// ---------- queries ----------

// Employee: employee(id: ID!)
func (r *Resolver) Employee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return nil, err
	}
	e, err := r.employees.GetByID(id)
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return e, nil
}

// Employees: employees(page, pageSize, search, departmentId, minSalary, maxSalary, sortBy, sortOrder)
func (r *Resolver) Employees(p gql.ResolveParams) (interface{}, error) {
	filter, err := filterFromArgs(p.Args)
	if err != nil {
		return nil, err
	}
	page, err := r.employees.List(filter)
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return page, nil
}

// ---------- nested field: Employee -> department ----------

// The source of a nested resolver can be a value or a pointer.
func asEmployee(src interface{}) (*model.Employee, bool) {
	switch v := src.(type) {
	case model.Employee:
		return &v, true
	case *model.Employee:
		return v, v != nil
	}
	return nil, false
}

// EmployeeDepartment resolves  employee { department { ... } }.
// The repository already loaded the department with Preload (one query for the whole list),
// so normally no extra query happens here. The fallback only runs if it was not preloaded.
func (r *Resolver) EmployeeDepartment(p gql.ResolveParams) (interface{}, error) {
	e, ok := asEmployee(p.Source)
	if !ok {
		return nil, newError("INTERNAL_SERVER_ERROR", "something went wrong")
	}
	if e.Department != nil {
		return e.Department, nil
	}
	d, err := r.departments.GetByID(e.DepartmentID)
	if err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return d, nil
}

// ---------- mutations ----------

func (r *Resolver) CreateEmployee(p gql.ResolveParams) (interface{}, error) {
	e, err := employeeFromInput(p.Args["input"].(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	if err := r.employees.Create(&e); err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return e, nil
}

func (r *Resolver) UpdateEmployee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return nil, err
	}
	e, err := employeeFromInput(p.Args["input"].(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	updated, err := r.employees.Update(id, &e)
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return updated, nil
}

func (r *Resolver) DeleteEmployee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return nil, err
	}
	if err := r.employees.Delete(id); err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return true, nil
}
