package graphql

import gql "github.com/graphql-go/graphql"

func newMutationType(
	server *Server,
	employeeType *gql.Object,
	employeeInput *gql.InputObject,
) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(employeeInput),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					employee, err := employeeFromInput(p.Args["input"])
					if err != nil {
						return nil, err
					}

					if err := server.service.Create(&employee); err != nil {
						return nil, err
					}

					return &employee, nil
				},
			},
			"updateEmployee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
					"input": &gql.ArgumentConfig{
						Type: gql.NewNonNull(employeeInput),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return nil, err
					}

					employee, err := employeeFromInput(p.Args["input"])
					if err != nil {
						return nil, err
					}

					return server.service.Update(id, employee)
				},
			},
			"deleteEmployee": &gql.Field{
				Type: gql.NewNonNull(gql.Boolean),
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return nil, err
					}

					if err := server.service.Delete(id); err != nil {
						return nil, err
					}

					return true, nil
				},
			},
		},
	})
}
