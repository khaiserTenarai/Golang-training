package resolver

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"order-management/services/product/graphql/schema"
	"order-management/services/product/model"
	"order-management/services/product/service"

	"github.com/graphql-go/graphql"
)

// Resolver connects GraphQL operations to the Product Service.
type Resolver struct {
	Service service.ProductService
}

// New creates a GraphQL schema with product queries and mutations.
func New(productService service.ProductService) (graphql.Schema, error) {
	r := &Resolver{Service: productService}

	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"products": &graphql.Field{
				Type: graphql.NewList(schema.ProductType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return r.Service.GetAll(context.Background())
				},
			},
			"product": &graphql.Field{
				Type: schema.ProductType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.ParseUint(fmt.Sprint(p.Args["id"]), 10, 64)
					if err != nil || id == 0 {
						return nil, fmt.Errorf("invalid product ID")
					}
					return r.Service.GetByID(context.Background(), uint(id))
				},
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createProduct": &graphql.Field{
				Type: schema.ProductType,
				Args: graphql.FieldConfigArgument{
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(schema.CreateProductInput)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := p.Args["input"].(map[string]interface{})
					product := &model.Product{
						Name:        fmt.Sprint(input["name"]),
						Description: optionalString(input["description"]),
						Price:       input["price"].(float64),
						Stock:       input["stock"].(int),
					}
					return r.Service.Create(context.Background(), product)
				},
			},
			"updateProduct": &graphql.Field{
				Type: schema.ProductType,
				Args: graphql.FieldConfigArgument{
					"id":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
					"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(schema.UpdateProductInput)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.ParseUint(fmt.Sprint(p.Args["id"]), 10, 64)
					if err != nil || id == 0 {
						return nil, fmt.Errorf("invalid product ID")
					}

					input := p.Args["input"].(map[string]interface{})
					existing, err := r.Service.GetByID(context.Background(), uint(id))
					if err != nil {
						return nil, err
					}

					if value, ok := input["name"].(string); ok {
						existing.Name = value
					}
					if value, ok := input["description"].(string); ok {
						existing.Description = value
					}
					if value, ok := input["price"].(float64); ok {
						existing.Price = value
					}
					if value, ok := input["stock"].(int); ok {
						existing.Stock = value
					}

					return r.Service.Update(context.Background(), uint(id), existing)
				},
			},
			"deleteProduct": &graphql.Field{
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.ParseUint(fmt.Sprint(p.Args["id"]), 10, 64)
					if err != nil || id == 0 {
						return false, fmt.Errorf("invalid product ID")
					}
					if err := r.Service.Delete(context.Background(), uint(id)); err != nil {
						return false, err
					}
					return true, nil
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})
}

func optionalString(value interface{}) string {
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}

// FormatProductTime formats a timestamp for GraphQL output.
func FormatProductTime(value time.Time) string {
	return value.Format(time.RFC3339)
}
