package main

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
)

func main() {
	// Authentication token checked from context
	ctx := context.WithValue(context.Background(), "userToken", "valid-jwt-token")

	resolve := func(p graphql.ResolveParams) (interface{}, error) {
		token, ok := p.Context.Value("userToken").(string)
		if !ok || token == "" {
			return nil, fmt.Errorf("unauthenticated request")
		}
		return "Authenticated Data Access Granted", nil
	}

	res, err := resolve(graphql.ResolveParams{Context: ctx})
	fmt.Printf("Auth result: %v (err: %v)
", res, err)
}
