package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

// Shared Model
type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Shared Service Interface & Implementation
type EmployeeService interface {
	GetEmployee() Employee
}

type employeeService struct{}

func (s *employeeService) GetEmployee() Employee {
	return Employee{ID: "e101", Name: "Vimal"}
}

// REST Handler sharing the Service
func restGetEmployee(svc EmployeeService) gin.HandlerFunc {
	return func(c *gin.Context) {
		emp := svc.GetEmployee()
		c.JSON(200, emp)
	}
}

func main() {
	// Initialize Shared Service
	svc := &employeeService{}

	// Setup GraphQL Schema sharing the same Service
	employeeType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Employee",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.String},
			"name": &graphql.Field{Type: graphql.String},
		},
	})

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"employee": &graphql.Field{
				Type: employeeType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return svc.GetEmployee(), nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: rootQuery})
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	// 1. REST Endpoint
	r.GET("/api/employee", restGetEmployee(svc))

	// 2. GraphQL Endpoint
	h := handler.New(&handler.Config{Schema: &schema, GraphiQL: true})
	r.Any("/graphql", gin.WrapH(h))

	fmt.Println("REST:    http://localhost:8080/api/employee")
	fmt.Println("GraphQL: http://localhost:8080/graphql")
	r.Run(":8080")
}