package graphql

import (
	"employee-management/model"
	"fmt"
	gql "github.com/graphql-go/graphql"
	"strconv"
	"strings"
)

// ---------- helpers ----------

// parseID converts a GraphQL ID (always a string) into a uint.
func parseID(v interface{}) (uint, error) {
	id, err := strconv.ParseUint(fmt.Sprint(v), 10, 32)
	if err != nil || id == 0 {
		return 0, newError("BAD_USER_INPUT", "invalid id")
	}
	return uint(id), nil
}
func str(v interface{}) string {
	s, _ := v.(string)
	return strings.TrimSpace(s)
}
func employeeFromInput(m map[string]interface{}) (model.Employee, error) {
	deptID, err := parseID(m["departmentId"])
	if err != nil {
		return model.Employee{}, err
	}
	salary, _ := m["salary"].(float64)
	return model.Employee{Name: str(m["name"]), Email: str(m["email"]), Salary: salary, DepartmentID: deptID}, nil
}

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
func asDepartment(src interface{}) (*model.Department, bool) {
	switch v := src.(type) {
	case model.Department:
		return &v, true
	case *model.Department:
		return v, v != nil
	}
	return nil, false
}

// ---------- query resolvers ----------

func (x *Server) resolveEmployee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return nil, err
	}
	e, err := x.employees.GetByID(id)
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return e, nil
}
func (x *Server) resolveEmployees(p gql.ResolveParams) (interface{}, error) {
	f := model.EmployeeFilter{}
	f.Limit, _ = p.Args["limit"].(int)
	f.Offset, _ = p.Args["offset"].(int)
	if f.Limit < 1 || f.Limit > 100 {
		return nil, newError("BAD_USER_INPUT", "limit must be between 1 and 100")
	}
	if f.Offset < 0 {
		return nil, newError("BAD_USER_INPUT", "offset cannot be negative")
	}
	if s, ok := p.Args["search"].(string); ok {
		f.Search = strings.TrimSpace(s)
	}
	if v, ok := p.Args["departmentId"]; ok && v != nil {
		id, err := parseID(v)
		if err != nil {
			return nil, err
		}
		f.DepartmentID = id
	}
	list, err := x.employees.GetAll(f)
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return list, nil
}
func (x *Server) resolveDepartment(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return nil, err
	}
	d, err := x.departments.GetByID(id)
	if err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return d, nil
}
func (x *Server) resolveDepartments(p gql.ResolveParams) (interface{}, error) {
	list, err := x.departments.GetAll()
	if err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return list, nil
}

// ---------- nested field resolvers ----------

// Employee.department: already loaded by the repository (Preload), so no extra query is normally needed.
func (x *Server) resolveEmployeeDepartment(p gql.ResolveParams) (interface{}, error) {
	e, ok := asEmployee(p.Source)
	if !ok {
		return nil, newError("INTERNAL_SERVER_ERROR", "something went wrong")
	}
	if e.Department != nil {
		return e.Department, nil
	}
	d, err := x.departments.GetByID(e.DepartmentID)
	if err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return d, nil
}

// Department.employees
func (x *Server) resolveDepartmentEmployees(p gql.ResolveParams) (interface{}, error) {
	d, ok := asDepartment(p.Source)
	if !ok {
		return nil, newError("INTERNAL_SERVER_ERROR", "something went wrong")
	}
	list, err := x.employees.GetAll(model.EmployeeFilter{DepartmentID: d.ID})
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return list, nil
}

// ---------- mutation resolvers ----------

func (x *Server) resolveCreateDepartment(p gql.ResolveParams) (interface{}, error) {
	m := p.Args["input"].(map[string]interface{})
	d := model.Department{Name: str(m["name"]), Description: str(m["description"])}
	if err := x.departments.Create(&d); err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return d, nil
}
func (x *Server) resolveCreateEmployee(p gql.ResolveParams) (interface{}, error) {
	e, err := employeeFromInput(p.Args["input"].(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	if err := x.employees.Create(&e); err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return e, nil
}
func (x *Server) resolveUpdateEmployee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return nil, err
	}
	e, err := employeeFromInput(p.Args["input"].(map[string]interface{}))
	if err != nil {
		return nil, err
	}
	updated, err := x.employees.Update(id, &e)
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return updated, nil
}
func (x *Server) resolveDeleteEmployee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return false, err
	}
	if err := x.employees.Delete(id); err != nil {
		return false, toGraphQLError(err, "employee")
	}
	return true, nil
}
