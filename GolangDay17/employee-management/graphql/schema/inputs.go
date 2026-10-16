package schema

import gql "github.com/graphql-go/graphql"

// inputTypes are the shapes of data that clients SEND to the API.
type inputTypes struct {
	EmployeeInput   *gql.InputObject
	DepartmentInput *gql.InputObject
	SortField       *gql.Enum
	SortOrder       *gql.Enum
}

func newInputTypes() *inputTypes {
	employeeInput := gql.NewInputObject(gql.InputObjectConfig{Name: "EmployeeInput", Fields: gql.InputObjectConfigFieldMap{
		"name":         &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"email":        &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"salary":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
		"departmentId": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.ID)},
	}})

	departmentInput := gql.NewInputObject(gql.InputObjectConfig{Name: "DepartmentInput", Fields: gql.InputObjectConfigFieldMap{
		"name":        &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"description": &gql.InputObjectFieldConfig{Type: gql.String},
	}})

	// An ENUM only allows a fixed list of values. Clients write ID, NAME ...
	// and the resolver receives the "Value" (id, name ...). Nobody can sort by a column that is not listed.
	sortField := gql.NewEnum(gql.EnumConfig{Name: "EmployeeSortField", Values: gql.EnumValueConfigMap{
		"ID":         &gql.EnumValueConfig{Value: "id"},
		"NAME":       &gql.EnumValueConfig{Value: "name"},
		"EMAIL":      &gql.EnumValueConfig{Value: "email"},
		"SALARY":     &gql.EnumValueConfig{Value: "salary"},
		"CREATED_AT": &gql.EnumValueConfig{Value: "createdAt"},
	}})
	sortOrder := gql.NewEnum(gql.EnumConfig{Name: "SortOrder", Values: gql.EnumValueConfigMap{
		"ASC":  &gql.EnumValueConfig{Value: "asc"},
		"DESC": &gql.EnumValueConfig{Value: "desc"},
	}})

	return &inputTypes{
		EmployeeInput:   employeeInput,
		DepartmentInput: departmentInput,
		SortField:       sortField,
		SortOrder:       sortOrder,
	}
}
