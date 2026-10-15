package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

// Service Interface
type EmployeeService interface {
	GetEmployeeByID(id int) (interface{}, error)
}

// Resolver Handler with Service Injected
type Resolver struct {
	Service EmployeeService
}

// Resolver Method using Injected Service
func (r *Resolver) ResolveGetEmployee(p graphql.ResolveParams) (interface{}, error) {
	id := p.Args["id"].(int)
	return r.Service.GetEmployeeByID(id)
}

func main() {
	fmt.Println("Service Injected into Resolver")
}