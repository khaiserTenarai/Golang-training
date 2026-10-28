package graphql

import (
	"reflect"

	"employee-management/model"

	gql "github.com/graphql-go/graphql"
)

var DepartmentType = gql.NewObject(gql.ObjectConfig{
	Name: "Department",
	Fields: gql.Fields{
		"name": &gql.Field{
			Type: gql.String,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				if name, ok := p.Source.(string); ok {
					return name, nil
				}
				return nil, nil
			},
		},
	},
})

var EmployeeType = gql.NewObject(gql.ObjectConfig{
	Name: "Employee",
	Fields: gql.Fields{
		"id":    &gql.Field{Type: gql.Int},
		"name":  &gql.Field{Type: gql.String},
		"email": &gql.Field{Type: gql.String},
		"department": &gql.Field{
			Type: DepartmentType,
			Resolve: func(p gql.ResolveParams) (interface{}, error) {
				// The reference model stores Department as a string.
				// Pass that string as the nested Department object's source.
				switch employee := p.Source.(type) {
				case model.Employee:
					return employee.Department, nil
				case *model.Employee:
					return employee.Department, nil
				case map[string]interface{}:
					return employee["department"], nil
				default:
					value := reflect.ValueOf(p.Source)
					if value.IsValid() && value.Kind() == reflect.Ptr && !value.IsNil() {
						value = value.Elem()
					}
					if value.IsValid() && value.Kind() == reflect.Struct {
						field := value.FieldByName("Department")
						if field.IsValid() && field.Kind() == reflect.String {
							return field.String(), nil
						}
					}
					return nil, nil
				}
			},
		},
		"salary": &gql.Field{Type: gql.Float},
	},
})

var EmployeeInputType = gql.NewInputObject(gql.InputObjectConfig{
	Name: "EmployeeInput",
	Fields: gql.InputObjectConfigFieldMap{
		"name":       &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"email":      &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"department": &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.String)},
		"salary":     &gql.InputObjectFieldConfig{Type: gql.NewNonNull(gql.Float)},
	},
})
