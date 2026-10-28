package graphql

import (
	"day17-employee-platform/model"

	gql "github.com/graphql-go/graphql"
)

func buildQuery(s *Server, employeeType, departmentType *gql.Object) *gql.Object {
	pageType := gql.NewObject(gql.ObjectConfig{
		Name: "EmployeePage",
		Fields: gql.Fields{
			"items": {
				Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(employeeType))),
			},
			"total": {
				Type: gql.NewNonNull(gql.Int),
			},
			"page": {
				Type: gql.NewNonNull(gql.Int),
			},
			"pageSize": {
				Type: gql.NewNonNull(gql.Int),
			},
			"totalPages": {
				Type: gql.NewNonNull(gql.Int),
			},
		},
	})

	return gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employees": {
				Type: gql.NewNonNull(pageType),
				Args: gql.FieldConfigArgument{
					"page":      {Type: gql.Int},
					"pageSize":  {Type: gql.Int},
					"name":      {Type: gql.String},
					"sortBy":    {Type: gql.String},
					"sortOrder": {Type: gql.String},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					page := 1
					pageSize := 5

					if value, ok := p.Args["page"].(int); ok && value > 0 {
						page = value
					}
					if value, ok := p.Args["pageSize"].(int); ok && value > 0 {
						pageSize = value
					}

					name, _ := p.Args["name"].(string)
					sortBy, _ := p.Args["sortBy"].(string)
					sortOrder, _ := p.Args["sortOrder"].(string)

					items, total, totalPages, err := s.EmployeeService.GetAll(
						model.EmployeeFilter{Name: name},
						model.SortOption{Field: sortBy, Order: sortOrder},
						page,
						pageSize,
					)
					if err != nil {
						return nil, err
					}

					result := make([]*model.Employee, 0, len(items))
					for i := range items {
						result = append(result, &items[i])
					}

					return map[string]interface{}{
						"items":      result,
						"total":      int(total),
						"page":       page,
						"pageSize":   pageSize,
						"totalPages": totalPages,
					}, nil
				},
			},
			"employee": {
				Type: employeeType,
				Args: gql.FieldConfigArgument{
					"id": {Type: gql.NewNonNull(gql.Int)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := uint(p.Args["id"].(int))
					return s.EmployeeService.GetByID(id)
				},
			},
			"department": {
				Type: departmentType,
				Args: gql.FieldConfigArgument{
					"id": {Type: gql.NewNonNull(gql.Int)},
				},
				Resolve: func(p gql.ResolveParams) (interface{}, error) {
					id := uint(p.Args["id"].(int))
					return s.DepartmentService.GetByID(id)
				},
			},
		},
	})
}
