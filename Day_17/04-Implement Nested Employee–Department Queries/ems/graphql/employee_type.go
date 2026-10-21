package graphql

import (
	"employee-management/model"

	gql "github.com/graphql-go/graphql"
)

func newDepartmentType() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Department",
		Fields: gql.Fields{
			"name": &gql.Field{
				Type: gql.NewNonNull(gql.String),
			},
		},
	})
}

func newEmployeeType() *gql.Object {
	departmentType := newDepartmentType()

	return gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id": &gql.Field{
				Type: gql.NewNonNull(gql.ID),
			},
			"name": &gql.Field{
				Type: gql.NewNonNull(gql.String),
			},
			"email": &gql.Field{
				Type: gql.NewNonNull(gql.String),
			},
			"department": &gql.Field{
				Type: gql.NewNonNull(departmentType),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					var department string

					switch employee := p.Source.(type) {
					case model.Employee:
						department = employee.Department
					case *model.Employee:
						if employee == nil {
							return nil, nil
						}
						department = employee.Department
					default:
						return nil, nil
					}

					return map[string]interface{}{
						"name": department,
					}, nil
				},
			},
			"salary": &gql.Field{
				Type: gql.NewNonNull(gql.Float),
			},
		},
	})
}
