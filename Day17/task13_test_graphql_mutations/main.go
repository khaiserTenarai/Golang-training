package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

func main() {
	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"add": &graphql.Field{
				Type: graphql.String,
				Args: graphql.FieldConfigArgument{
					"msg": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Added: " + p.Args["msg"].(string), nil
				},
			},
		},
	})

	schema, _ := graphql.NewSchema(graphql.SchemaConfig{Mutation: mutationType})
	res := graphql.Do(graphql.Params{Schema: schema, RequestString: `mutation { add(msg: "Hello") }`})

	fmt.Printf("Mutation Test Result: %+v
", res.Data)
}
