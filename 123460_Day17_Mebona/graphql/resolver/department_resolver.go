package resolver

import (
	"employee-management/model"
	gql "github.com/graphql-go/graphql"
)

func (r *Resolver) Department(p gql.ResolveParams) (interface{}, error) {
	id, err := parseID(p.Args["id"])
	if err != nil {
		return nil, err
	}
	d, err := r.departments.GetByID(id)
	if err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return d, nil
}

func (r *Resolver) Departments(p gql.ResolveParams) (interface{}, error) {
	list, err := r.departments.GetAll()
	if err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return list, nil
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

// DepartmentEmployees resolves  department { employees { ... } }
func (r *Resolver) DepartmentEmployees(p gql.ResolveParams) (interface{}, error) {
	d, ok := asDepartment(p.Source)
	if !ok {
		return nil, newError("INTERNAL_SERVER_ERROR", "something went wrong")
	}
	list, err := r.employees.ListByDepartment(d.ID)
	if err != nil {
		return nil, toGraphQLError(err, "employee")
	}
	return list, nil
}

func (r *Resolver) CreateDepartment(p gql.ResolveParams) (interface{}, error) {
	m := p.Args["input"].(map[string]interface{})
	d := model.Department{Name: str(m["name"]), Description: str(m["description"])}
	if err := r.departments.Create(&d); err != nil {
		return nil, toGraphQLError(err, "department")
	}
	return d, nil
}
