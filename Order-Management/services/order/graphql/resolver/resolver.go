package resolver

import (
	"context"
	"fmt"
	"strconv"

	"order-management/services/order/graphql/schema"
	"order-management/services/order/service"

	"github.com/graphql-go/graphql"
)

func New(orderService service.OrderService) (graphql.Schema, error) {
	query := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"orders": &graphql.Field{
				Type: graphql.NewList(schema.OrderType),
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return orderService.GetAll(p.Context)
				},
			},
			"order": &graphql.Field{
				Type: schema.OrderType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, err := strconv.ParseUint(fmt.Sprint(p.Args["id"]), 10, 64)
					if err != nil || id == 0 {
						return nil, fmt.Errorf("invalid order ID")
					}
					return orderService.GetByID(
						p.Context,
						uint(id),
					)
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{Query: query})
}

// Ensure request contexts can be used by GraphQL resolvers.
var _ = context.Background
