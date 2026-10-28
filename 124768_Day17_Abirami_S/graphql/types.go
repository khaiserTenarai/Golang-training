package graphql

import (
	"day17-employee-platform/model"

	gql "github.com/graphql-go/graphql"
)

func departmentType() *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Department",
		Fields: gql.Fields{
			"id": {
				Type: gql.NewNonNull(gql.Int),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					department, ok := p.Source.(*model.Department)
					if !ok {
						return nil, nil
					}
					return int(department.ID), nil
				},
			},
			"name": {
				Type: gql.NewNonNull(gql.String),
			},
		},
	})
}

func employeeType(_ *Server, departmentType *gql.Object) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Employee",
		Fields: gql.Fields{
			"id": {
				Type: gql.NewNonNull(gql.Int),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					employee, ok := p.Source.(*model.Employee)
					if !ok {
						return nil, nil
					}
					return int(employee.ID), nil
				},
			},
			"name": {
				Type: gql.NewNonNull(gql.String),
			},
			"email": {
				Type: gql.NewNonNull(gql.String),
			},
			"salary": {
				Type: gql.NewNonNull(gql.Float),
			},
			"department": {
				Type: departmentType,
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					employee, ok := p.Source.(*model.Employee)
					if !ok {
						return nil, nil
					}

					// Department is already loaded by GORM Preload("Department").
					return employee.Department, nil
				},
			},
		},
	})
}
