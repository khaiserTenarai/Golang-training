package resolver

import (
	"employee-management/auth"
	gql "github.com/graphql-go/graphql"
)

// Login: login(username, password) -> token. This one is public (no token needed).
func (r *Resolver) Login(p gql.ResolveParams) (interface{}, error) {
	password, _ := p.Args["password"].(string) // passwords are NOT trimmed
	result, err := r.users.Login(str(p.Args["username"]), password)
	if err != nil {
		return nil, toGraphQLError(err, "user")
	}
	return result, nil
}

// Me: me -> the user stored in the token.
func (r *Resolver) Me(p gql.ResolveParams) (interface{}, error) {
	return auth.UserFrom(p.Context), nil
}
