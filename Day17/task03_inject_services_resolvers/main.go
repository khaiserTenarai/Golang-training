package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

type EmployeeService struct{}

func (s *EmployeeService) FetchByID(id string) map[string]string {
	return map[string]string{"id": id, "name": "Alice"}
}

type Resolver struct {
	Service *EmployeeService // Service Injected
}

func (r *Resolver) ResolveEmployee(p graphql.ResolveParams) (interface{}, error) {
	return r.Service.FetchByID("emp-1"), nil
}

func main() {
	svc := &EmployeeService{}
	r := &Resolver{Service: svc}
	res, _ := r.ResolveEmployee(graphql.ResolveParams{})
	fmt.Printf("Injected Service Result: %+v
", res)
}
