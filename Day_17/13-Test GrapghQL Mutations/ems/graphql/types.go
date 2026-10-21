package graphql

import (
	"employee-management/model"

	gql "github.com/graphql-go/graphql"
)

func newDepartmentType() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Department",
		Fields: gql.Fields{
			"name": &gql.Field{Type: gql.NewNonNull(gql.String)},
		},
	})
}

func newEmployeeType() *gql.Object {
	departmentType := newDepartmentType()

	return gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id":     &gql.Field{Type: gql.NewNonNull(gql.ID)},
			"name":   &gql.Field{Type: gql.NewNonNull(gql.String)},
			"email":  &gql.Field{Type: gql.NewNonNull(gql.String)},
			"salary": &gql.Field{Type: gql.NewNonNull(gql.Float)},

			"department": &gql.Field{
				Type: gql.NewNonNull(departmentType),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					switch e := p.Source.(type) {
					case model.Employee:
						return map[string]interface{}{"name": e.Department}, nil
					case *model.Employee:
						if e == nil {
							return nil, nil
						}
						return map[string]interface{}{"name": e.Department}, nil
					default:
						return nil, nil
					}
				},
			},
		},
	})
}

func newEmployeeInputType() *gql.InputObject {
	return gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		},
	})
}

func newEmployeeFilterType() *gql.InputObject {
	return gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeFilter",
		Fields: gql.InputObjectConfigFieldMap{
			"name":       &gql.InputObjectFieldConfig{Type: gql.String},
			"department": &gql.InputObjectFieldConfig{Type: gql.String},
		},
	})
}

func newEmployeeSortType() *gql.InputObject {
	return gql.NewInputObject(gql.InputObjectConfig{
		Name: "EmployeeSort",
		Fields: gql.InputObjectConfigFieldMap{
			"field": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
			"direction": &gql.InputObjectFieldConfig{
				Type: gql.NewNonNull(gql.String),
			},
		},
	})
}

func newPaginationType() *gql.InputObject {
	return gql.NewInputObject(gql.InputObjectConfig{
		Name: "PaginationInput",
		Fields: gql.InputObjectConfigFieldMap{
			"page":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Int)},
			"pageSize": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Int)},
		},
	})
}

func newEmployeeListType(employeeType *gql.Object) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "EmployeeList",
		Fields: gql.Fields{
			"data": &gql.Field{
				Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(employeeType))),
			},
			"page":       &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"pageSize":   &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"total":      &gql.Field{Type: gql.NewNonNull(gql.Int)},
			"totalPages": &gql.Field{Type: gql.NewNonNull(gql.Int)},
		},
	})
}
