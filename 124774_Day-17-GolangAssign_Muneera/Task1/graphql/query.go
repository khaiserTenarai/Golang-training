package graphql

import (
	"fmt"
	"strconv"

	gql "github.com/graphql-go/graphql"

	"employee-management/service"
)

func NewQueryType(svc service.EmployeeService) *gql.Object {
	return gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employees": &gql.Field{
				Type: gql.NewNonNull(
					gql.NewList(gql.NewNonNull(EmployeeType)),
				),
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					return svc.GetAll()
				},
			},

			"employee": &gql.Field{
				Type: EmployeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{
						Type: gql.NewNonNull(gql.ID),
					},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					idText := fmt.Sprint(p.Args["id"])
					idValue, err := strconv.ParseUint(idText, 10, 32)
					if err != nil || idValue == 0 {
						return nil, fmt.Errorf("invalid employee ID")
					}

					return svc.GetByID(uint(idValue))
				},
			},
		},
	})
}
