package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

func main() {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"ping": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "pong", nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Query: queryType})
	res := graphql.Do(graphql.Params{Schema: schema, RequestString: "{ ping }"})

	fmt.Printf("Query Test Result: %+v
", res.Data)
}
