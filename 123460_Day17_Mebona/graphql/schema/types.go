package schema

import (
	"employee-management/graphql/resolver"
	gql "github.com/graphql-go/graphql"
)

// objectTypes are the "things" the API can return.
type objectTypes struct {
	Department   *gql.Object
	Employee     *gql.Object
	EmployeePage *gql.Object
	AuthPayload  *gql.Object
	CurrentUser  *gql.Object
}

func newObjectTypes(r *resolver.Resolver) *objectTypes {
	department := gql.NewObject(gql.ObjectConfig{Name: "Department", Description: "A company department", Fields: gql.Fields{
		"id":          &gql.Field{Type: gql.NewNonNull(gql.ID)},
		"name":        &gql.Field{Type: gql.NewNonNull(gql.String)},
		"description": &gql.Field{Type: gql.String},
	}})

	employee := gql.NewObject(gql.ObjectConfig{Name: "Employee", Description: "A person working in a department", Fields: gql.Fields{
		"id":        &gql.Field{Type: gql.NewNonNull(gql.ID)},
		"name":      &gql.Field{Type: gql.NewNonNull(gql.String)},
		"email":     &gql.Field{Type: gql.NewNonNull(gql.String)},
		"salary":    &gql.Field{Type: gql.NewNonNull(gql.Float)},
		"createdAt": &gql.Field{Type: gql.NewNonNull(gql.DateTime)},
		// nested query: Employee -> Department
		"department": &gql.Field{Type: gql.NewNonNull(department), Resolve: r.EmployeeDepartment},
	}})

	// Department -> employees is added afterwards because the two types point to each other.
	department.AddFieldConfig("employees", &gql.Field{
		Type:    gql.NewNonNull(gql.NewList(gql.NewNonNull(employee))),
		Resolve: r.DepartmentEmployees,
	})

	// One page of employees (pagination result)
	employeePage := gql.NewObject(gql.ObjectConfig{Name: "EmployeePage", Description: "One page of employees", Fields: gql.Fields{
		"items":      &gql.Field{Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(employee)))},
		"totalCount": &gql.Field{Type: gql.NewNonNull(gql.Int), Description: "All matching employees, not only this page"},
		"page":       &gql.Field{Type: gql.NewNonNull(gql.Int)},
		"pageSize":   &gql.Field{Type: gql.NewNonNull(gql.Int)},
		"totalPages": &gql.Field{Type: gql.NewNonNull(gql.Int)},
		"hasNext":    &gql.Field{Type: gql.NewNonNull(gql.Boolean)},
	}})

	authPayload := gql.NewObject(gql.ObjectConfig{Name: "AuthPayload", Description: "Result of a successful login", Fields: gql.Fields{
		"token":    &gql.Field{Type: gql.NewNonNull(gql.String)},
		"username": &gql.Field{Type: gql.NewNonNull(gql.String)},
		"role":     &gql.Field{Type: gql.NewNonNull(gql.String)},
	}})

	currentUser := gql.NewObject(gql.ObjectConfig{Name: "CurrentUser", Description: "The logged-in user", Fields: gql.Fields{
		"username": &gql.Field{Type: gql.NewNonNull(gql.String)},
		"role":     &gql.Field{Type: gql.NewNonNull(gql.String)},
	}})

	return &objectTypes{
		Department:   department,
		Employee:     employee,
		EmployeePage: employeePage,
		AuthPayload:  authPayload,
		CurrentUser:  currentUser,
	}
}
