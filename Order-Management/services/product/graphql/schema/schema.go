package schema

import "github.com/graphql-go/graphql"

// ProductType defines the GraphQL representation of a product.
var ProductType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"name":        &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"price":       &graphql.Field{Type: graphql.Float},
		"stock":       &graphql.Field{Type: graphql.Int},
		"createdAt":   &graphql.Field{Type: graphql.String},
		"updatedAt":   &graphql.Field{Type: graphql.String},
	},
})

// CreateProductInput defines fields accepted when creating a product.
var CreateProductInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "CreateProductInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"price":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
		"stock":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
	},
})

// UpdateProductInput defines fields accepted when updating a product.
var UpdateProductInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UpdateProductInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":        &graphql.InputObjectFieldConfig{Type: graphql.String},
		"description": &graphql.InputObjectFieldConfig{Type: graphql.String},
		"price":       &graphql.InputObjectFieldConfig{Type: graphql.Float},
		"stock":       &graphql.InputObjectFieldConfig{Type: graphql.Int},
	},
})
