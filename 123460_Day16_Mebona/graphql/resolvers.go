package graphql

import (
	"strconv"
	"time"

	"employee-management/apperror"
	"employee-management/model"
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

// resolver holds the services. Each method is a GraphQL resolver:
// read arguments -> call the service -> return value or error.
type resolver struct {
	employees   service.EmployeeService
	departments service.DepartmentService
}

// result converts service errors (via apperror.Wrap) and avoids returning typed nils.
func result(v interface{}, err error) (interface{}, error) {
	if err != nil {
		return nil, apperror.Wrap(err)
	}
	return v, nil
}

// ---------------- Query ----------------

func (r *resolver) employee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID("id", p.Args["id"])
	if err != nil {
		return nil, err
	}
	return result(r.employees.Get(id))
}

func (r *resolver) employeeList(p gql.ResolveParams) (interface{}, error) {
	f := model.EmployeeFilter{Limit: intArg(p.Args, "limit", 10), Offset: intArg(p.Args, "offset", 0)}
	f.Search, _ = p.Args["search"].(string)

	if raw, ok := p.Args["departmentId"]; ok && raw != nil {
		id, err := parseID("departmentId", raw)
		if err != nil {
			return nil, err
		}
		f.DepartmentID = &id
	}
	if t, ok := p.Args["joinedAfter"].(time.Time); ok {
		f.JoinedAfter = &t
	}
	return result(r.employees.List(f))
}

func (r *resolver) department(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID("id", p.Args["id"])
	if err != nil {
		return nil, err
	}
	return result(r.departments.Get(id))
}

func (r *resolver) departmentList(p gql.ResolveParams) (interface{}, error) {
	return result(r.departments.List())
}

// ---------------- Nested fields ----------------

// Employee.department: already loaded by GORM Preload, otherwise looked up.
func (r *resolver) employeeDepartment(p gql.ResolveParams) (interface{}, error) {
	emp := p.Source.(*model.Employee)
	if emp.Department.ID != 0 {
		return &emp.Department, nil
	}
	return result(r.departments.Get(emp.DepartmentID))
}

// Department.employees
func (r *resolver) departmentEmployees(p gql.ResolveParams) (interface{}, error) {
	dept := p.Source.(*model.Department)
	return result(r.departments.Employees(dept.ID))
}

// Department.employeeCount
func (r *resolver) departmentEmployeeCount(p gql.ResolveParams) (interface{}, error) {
	dept := p.Source.(*model.Department)
	list, err := r.departments.Employees(dept.ID)
	if err != nil {
		return nil, apperror.Wrap(err)
	}
	return len(list), nil
}

// ---------------- Mutation ----------------

func (r *resolver) createEmployee(p gql.ResolveParams) (interface{}, error) {
	in, _ := p.Args["input"].(map[string]interface{})
	deptID, err := parseID("departmentId", in["departmentId"])
	if err != nil {
		return nil, err
	}
	name, _ := in["name"].(string)
	email, _ := in["email"].(string)
	salary, _ := in["salary"].(float64)

	create := model.CreateEmployeeInput{Name: name, Email: email, Salary: salary, DepartmentID: deptID}
	if v, ok := in["active"].(bool); ok { // the schema's default value (true) is applied for us
		create.Active = &v
	}
	return result(r.employees.Create(create))
}

func (r *resolver) updateEmployee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID("id", p.Args["id"])
	if err != nil {
		return nil, err
	}
	in, _ := p.Args["input"].(map[string]interface{})

	upd := model.UpdateEmployeeInput{}
	if v, ok := in["name"].(string); ok {
		upd.Name = &v
	}
	if v, ok := in["email"].(string); ok {
		upd.Email = &v
	}
	if v, ok := in["salary"].(float64); ok {
		upd.Salary = &v
	}
	if v, ok := in["active"].(bool); ok {
		upd.Active = &v
	}
	if raw, ok := in["departmentId"]; ok && raw != nil {
		deptID, err := parseID("departmentId", raw)
		if err != nil {
			return nil, err
		}
		upd.DepartmentID = &deptID
	}
	return result(r.employees.Update(id, upd))
}

func (r *resolver) deleteEmployee(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID("id", p.Args["id"])
	if err != nil {
		return nil, err
	}
	if err := r.employees.Delete(id); err != nil {
		return nil, apperror.Wrap(err)
	}
	return true, nil
}

// ---------------- helpers ----------------

// GraphQL's ID scalar arrives as a string; the database uses uint.
func parseID(field string, v interface{}) (uint, error) {
	s, _ := v.(string)
	id, err := strconv.ParseUint(s, 10, 32)
	if err != nil || id == 0 {
		return 0, apperror.BadInput(field, field+" must be a positive integer")
	}
	return uint(id), nil
}

func intArg(args map[string]interface{}, key string, def int) int {
	if v, ok := args[key].(int); ok {
		return v
	}
	return def
}
