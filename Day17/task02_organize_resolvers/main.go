package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

type EmployeeResolver struct{}

func (r *EmployeeResolver) GetEmployee(p graphql.ResolveParams) (interface{}, error) {
	return map[string]interface{}{"id": "101", "name": "Jane Doe"}, nil
}

func main() {
	resolver := &EmployeeResolver{}
	p := graphql.ResolveParams{}
	res, _ := resolver.GetEmployee(p)
	fmt.Printf("Resolver package structure output: %+v
", res)
}
