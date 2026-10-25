package graphql

import (
	"net/http"

	handler "github.com/graphql-go/handler"
)

// newGraphQLHTTPHandler creates the GraphQL execution handler.
// The custom GraphiQL interface is served separately at /graphiql.
func newGraphQLHTTPHandler(server *Server) http.Handler {
	return handler.New(&handler.Config{
		Schema:   server.Schema(),
		Pretty:   true,
		GraphiQL: false,
	})
}
