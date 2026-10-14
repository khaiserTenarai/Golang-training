package main

import (
	"context"
	"fmt"
	"github.com/graphql-go/graphql"
)

func main() {
	ctx := context.WithValue(context.Background(), "userRole", "USER")

	resolveAdminOnly := func(p graphql.ResolveParams) (interface{}, error) {
		role, _ := p.Context.Value("userRole").(string)
		if role != "ADMIN" {
			return nil, fmt.Errorf("forbidden: required role ADMIN, got %s", role)
		}
		return "Sensitive Admin Data", nil
	}

	_, err := resolveAdminOnly(graphql.ResolveParams{Context: ctx})
	fmt.Printf("Authorization Error Caught: %v
", err)
}
