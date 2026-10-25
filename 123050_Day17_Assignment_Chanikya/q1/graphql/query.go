package graphql

import gql "github.com/graphql-go/graphql"

func newQueryType(server *Server, employeeType *gql.Object) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employees": &gql.Field{
				Type: gql.NewNonNull(
					gql.NewList(gql.NewNonNull(employeeType)),
				),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return server.service.GetAll()
				},
			},
			"employee": &gql.Field{
				Type: employeeType,
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

					return server.service.GetByID(id)
				},
			},
		},
	})
}
