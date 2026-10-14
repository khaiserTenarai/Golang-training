package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

// Department Resolvers
func resolveGetDepartment(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil // Fetches department data
}

func resolveCreateDepartment(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil // Creates a new department
}

// Employee Resolvers
func resolveGetEmployee(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil // Fetches employee data
}

func resolveCreateEmployee(p graphql.ResolveParams) (interface{}, error) {
	return nil, nil // Creates a new employee
}

func main() {
	fmt.Println("Resolver Packages Organized")
}