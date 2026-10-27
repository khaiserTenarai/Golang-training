package schema

import "github.com/graphql-go/graphql"

var OrderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"customerId":  &graphql.Field{Type: graphql.Int},
		"productId":   &graphql.Field{Type: graphql.Int},
		"quantity":    &graphql.Field{Type: graphql.Int},
		"totalAmount": &graphql.Field{Type: graphql.Float},
		"status":      &graphql.Field{Type: graphql.String},
		"createdAt":   &graphql.Field{Type: graphql.String},
		"updatedAt":   &graphql.Field{Type: graphql.String},
	},
})
