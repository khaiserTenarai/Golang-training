package graphql

import (
	"fmt"

	"employee-management/model"

	gql "github.com/graphql-go/graphql"
)

func newQueryType(server *Server, employeeType *gql.Object) *gql.Object {
	filterType := newEmployeeFilterType()
	sortType := newEmployeeSortType()
	paginationType := newPaginationType()
	listType := newEmployeeListType(employeeType)

	return gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employees": &gql.Field{
				Type: gql.NewNonNull(listType),
				Args: gql.FieldConfigArgument{
					"filter":     &gql.ArgumentConfig{Type: filterType},
					"sort":       &gql.ArgumentConfig{Type: sortType},
					"pagination": &gql.ArgumentConfig{Type: paginationType},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					filter := model.EmployeeFilter{}
					if v, ok := p.Args["filter"].(map[string]interface{}); ok {
						filter.Name, _ = v["name"].(string)
						filter.Department, _ = v["department"].(string)
					}

					sort := model.EmployeeSort{Field: "id", Direction: "asc"}
					if v, ok := p.Args["sort"].(map[string]interface{}); ok {
						if x, ok := v["field"].(string); ok {
							sort.Field = x
						}
						if x, ok := v["direction"].(string); ok {
							sort.Direction = x
						}
					}

					page, pageSize := 1, 10
					if v, ok := p.Args["pagination"].(map[string]interface{}); ok {
						if x, ok := v["page"].(int); ok {
							page = x
						}
						if x, ok := v["pageSize"].(int); ok {
							pageSize = x
						}
					}

					result, err := server.service.GetAll(filter, sort, page, pageSize)
					if err != nil {
						return nil, fmt.Errorf("list employees: %w", err)
					}
					return result, nil
				},
			},

			"employee": &gql.Field{
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id, err := parseGraphQLID(p.Args["id"])
					if err != nil {
						return nil, err
					}
					return server.service.GetByID(id)
				},
			},
		},
	})
}

func parseGraphQLID(value interface{}) (uint, error) {
	var text string
	switch v := value.(type) {
	case string:
		text = v
	case int:
		text = fmt.Sprint(v)
	default:
		return 0, fmt.Errorf("invalid employee ID")
	}

	var id uint64
	if _, err := fmt.Sscan(text, &id); err != nil || id == 0 || fmt.Sprint(id) != text {
		return 0, fmt.Errorf("ID must be a positive integer")
	}
	return uint(id), nil
}
