package graphql

import (
	"employee-management/model"

	gql "github.com/graphql-go/graphql"
)

/*
Task 5: Optimize Nested Queries

The Employee model already contains Department as a string.
So, we don't need another database query to get the department name.

Example:
Employee.Department = "IT"

GraphQL response:
department: {
    name: "IT"
}
*/

/* Create the nested Department GraphQL type */
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

/* Create the Employee GraphQL type */
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

				/*
					Optimization:
					Read the department from the existing employee object.
					Do not call the service or repository again.
				*/
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					switch employee := p.Source.(type) {
					case model.Employee:
						return map[string]interface{}{
							"name": employee.Department,
						}, nil

					case *model.Employee:
						if employee == nil {
							return nil, nil
						}

						return map[string]interface{}{
							"name": employee.Department,
						}, nil
					}

					return nil, nil
				},
			},
			"salary": &gql.Field{
				Type: gql.NewNonNull(gql.Float),
			},
		},
	})
}
