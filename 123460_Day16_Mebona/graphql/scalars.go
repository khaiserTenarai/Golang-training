package graphql

import (
	"time"

	gql "github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
)

// DateTime is a custom scalar. Built-in scalars: ID, String, Int, Float, Boolean.
// On the wire it is an RFC3339 string ("2024-06-01T00:00:00Z"); in Go it is a time.Time.
var DateTime = gql.NewScalar(gql.ScalarConfig{
	Name:        "DateTime",
	Description: "RFC3339 timestamp, e.g. 2024-06-01T00:00:00Z",

	// Go -> JSON (output fields)
	Serialize: func(value interface{}) interface{} {
		switch v := value.(type) {
		case time.Time:
			return v.UTC().Format(time.RFC3339)
		case *time.Time:
			if v == nil {
				return nil
			}
			return v.UTC().Format(time.RFC3339)
		}
		return nil
	},

	// JSON "variables" -> Go
	ParseValue: func(value interface{}) interface{} {
		s, ok := value.(string)
		if !ok {
			return nil
		}
		t, err := time.Parse(time.RFC3339, s)
		if err != nil {
			return nil // nil = invalid value -> GraphQL reports a validation error
		}
		return t
	},

	// literal written inside the query text -> Go
	ParseLiteral: func(valueAST ast.Value) interface{} {
		sv, ok := valueAST.(*ast.StringValue)
		if !ok {
			return nil
		}
		t, err := time.Parse(time.RFC3339, sv.Value)
		if err != nil {
			return nil
		}
		return t
	},
})
