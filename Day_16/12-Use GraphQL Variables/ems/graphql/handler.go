package graphql

import (
	"net/http"

	graphqlhandler "github.com/graphql-go/handler"
)

func NewHTTPHandler(server *Server) http.Handler {
	return graphqlhandler.New(&graphqlhandler.Config{
		Schema:   server.Schema(),
		Pretty:   true,
		GraphiQL: true,
	})
}