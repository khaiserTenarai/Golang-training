// Package resolver documents the dependency-injection boundary for GraphQL.
// The executable graphql-go resolvers currently live together in ../graphql.go
// to keep the project beginner-friendly and avoid one file per resolver.
package resolver

import "employee-platform/service"

type Resolver struct { Service *service.Service }
func New(s *service.Service) *Resolver { return &Resolver{Service:s} }
