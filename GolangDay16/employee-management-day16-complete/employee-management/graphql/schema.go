package graphql

import (
	"employee-management/service"
	gql "github.com/graphql-go/graphql"
)

type Server struct {
	employees   service.EmployeeService
	departments service.DepartmentService
	schema      gql.Schema
}

func NewServer(e service.EmployeeService, d service.DepartmentService) (*Server, error) {
	x := &Server{employees: e, departments: d}
	sc, err := x.buildSchema()
	if err != nil {
		return nil, err
	}
	x.schema = sc
	return x, nil
}
func (x *Server) buildSchema() (gql.Schema, error) {
	// ---------- Object types ----------
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
		"department": &gql.Field{Type: gql.NewNonNull(department), Resolve: x.resolveEmployeeDepartment},
	}})
	// Department -> employees is added afterwards because the two types refer to each other.
	department.AddFieldConfig("employees", &gql.Field{
		Type:    gql.NewNonNull(gql.NewList(gql.NewNonNull(employee))),
		Resolve: x.resolveDepartmentEmployees,
	})

	// ---------- Input types ----------
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

	// ---------- Queries ----------
	idArg := gql.FieldConfigArgument{"id": &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)}}
	query := gql.NewObject(gql.ObjectConfig{Name: "Query", Fields: gql.Fields{
		"employee": &gql.Field{Description: "Get one employee by id", Type: employee, Args: idArg, Resolve: x.resolveEmployee},
		"employees": &gql.Field{
			Description: "List employees (optional filter and paging)",
			Type:        gql.NewNonNull(gql.NewList(gql.NewNonNull(employee))),
			Args: gql.FieldConfigArgument{
				"departmentId": &gql.ArgumentConfig{Type: gql.ID, Description: "Only employees of this department"},
				"search":       &gql.ArgumentConfig{Type: gql.String, Description: "Matches name or email"},
				"limit":        &gql.ArgumentConfig{Type: gql.Int, DefaultValue: 50, Description: "1 to 100"},
				"offset":       &gql.ArgumentConfig{Type: gql.Int, DefaultValue: 0},
			},
			Resolve: x.resolveEmployees,
		},
		"department":  &gql.Field{Description: "Get one department by id", Type: department, Args: idArg, Resolve: x.resolveDepartment},
		"departments": &gql.Field{Type: gql.NewNonNull(gql.NewList(gql.NewNonNull(department))), Resolve: x.resolveDepartments},
	}})

	// ---------- Mutations ----------
	mutation := gql.NewObject(gql.ObjectConfig{Name: "Mutation", Fields: gql.Fields{
		"createDepartment": &gql.Field{
			Type:    gql.NewNonNull(department),
			Args:    gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(departmentInput)}},
			Resolve: x.resolveCreateDepartment,
		},
		"createEmployee": &gql.Field{
			Type:    gql.NewNonNull(employee),
			Args:    gql.FieldConfigArgument{"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)}},
			Resolve: x.resolveCreateEmployee,
		},
		"updateEmployee": &gql.Field{
			Type: gql.NewNonNull(employee),
			Args: gql.FieldConfigArgument{
				"id":    &gql.ArgumentConfig{Type: gql.NewNonNull(gql.ID)},
				"input": &gql.ArgumentConfig{Type: gql.NewNonNull(employeeInput)},
			},
			Resolve: x.resolveUpdateEmployee,
		},
		"deleteEmployee": &gql.Field{Type: gql.NewNonNull(gql.Boolean), Args: idArg, Resolve: x.resolveDeleteEmployee},
	}})
	return gql.NewSchema(gql.SchemaConfig{Query: query, Mutation: mutation})
}
