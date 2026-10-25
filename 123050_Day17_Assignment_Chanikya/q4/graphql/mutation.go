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
					return server.resolvers.Create(p.Args["input"])
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
					return server.resolvers.Update(
						p.Args["id"],
						p.Args["input"],
					)
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
					return server.resolvers.Delete(p.Args["id"])
				},
			},
		},
	})
}
