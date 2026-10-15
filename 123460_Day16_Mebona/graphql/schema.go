package graphql

import (
	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

type Server struct {
	schema gql.Schema
}

func NewServer(emp service.EmployeeService, dept service.DepartmentService) (*Server, error) {
	sc, err := buildSchema(&resolver{employees: emp, departments: dept})
	if err != nil {
		return nil, err
	}
	return &Server{schema: sc}, nil
}

// buildSchema creates the complete schema (code-first). schema.graphql in the
// project root shows the same schema in SDL form.
func buildSchema(r *resolver) (gql.Schema, error) {
	nonNull := gql.NewNonNull

	// ---------- Object types ----------
	departmentType := gql.NewObject(gql.ObjectConfig{
		Name:        "Department",
		Description: "A department of the company",
		Fields: gql.Fields{
			"id":       {Type: nonNull(gql.ID)},
			"name":     {Type: nonNull(gql.String)},
			"location": {Type: nonNull(gql.String)},
		},
	})

	employeeType := gql.NewObject(gql.ObjectConfig{
		Name:        "Employee",
		Description: "A person working in the company",
		Fields: gql.Fields{
			"id":       {Type: nonNull(gql.ID)},      // ID scalar
			"name":     {Type: nonNull(gql.String)},  // String scalar
			"email":    {Type: nonNull(gql.String)},
			"salary":   {Type: nonNull(gql.Float)},   // Float scalar
			"active":   {Type: nonNull(gql.Boolean)}, // Boolean scalar
			"joinedAt": {Type: nonNull(DateTime)},    // custom scalar
		},
	})

	// The two types reference each other, so the links are added once both exist.
	employeeType.AddFieldConfig("department", &gql.Field{
		Type:    nonNull(departmentType),
		Resolve: r.employeeDepartment,
	})
	departmentType.AddFieldConfig("employees", &gql.Field{
		Type:    nonNull(gql.NewList(nonNull(employeeType))),
		Resolve: r.departmentEmployees,
	})
	departmentType.AddFieldConfig("employeeCount", &gql.Field{
		Type:    nonNull(gql.Int), // Int scalar
		Resolve: r.departmentEmployeeCount,
	})

	// ---------- Input types ----------
	createInput := gql.NewInputObject(gql.InputObjectConfig{
		Name: "CreateEmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":         {Type: nonNull(gql.String)},
			"email":        {Type: nonNull(gql.String)},
			"salary":       {Type: nonNull(gql.Float)},
			"departmentId": {Type: nonNull(gql.ID)},
			"active":       {Type: gql.Boolean, DefaultValue: true},
		},
	})

	// Every field optional -> partial update.
	updateInput := gql.NewInputObject(gql.InputObjectConfig{
		Name: "UpdateEmployeeInput",
		Fields: gql.InputObjectConfigFieldMap{
			"name":         {Type: gql.String},
			"email":        {Type: gql.String},
			"salary":       {Type: gql.Float},
			"departmentId": {Type: gql.ID},
			"active":       {Type: gql.Boolean},
		},
	})

	// ---------- Queries (read) ----------
	query := gql.NewObject(gql.ObjectConfig{
		Name: "Query",
		Fields: gql.Fields{
			"employee": {
				Type:    employeeType, // nullable: null + error when not found
				Args:    gql.FieldConfigArgument{"id": {Type: nonNull(gql.ID)}},
				Resolve: r.employee,
			},
			"employees": {
				Type: nonNull(gql.NewList(nonNull(employeeType))),
				Args: gql.FieldConfigArgument{
					"departmentId": {Type: gql.ID, Description: "only employees of this department"},
					"search":       {Type: gql.String, Description: "matches name or email"},
					"joinedAfter":  {Type: DateTime, Description: "only employees who joined after this time"},
					"limit":        {Type: gql.Int, DefaultValue: 10, Description: "1..100"},
					"offset":       {Type: gql.Int, DefaultValue: 0},
				},
				Resolve: r.employeeList,
			},
			"department": {
				Type:    departmentType,
				Args:    gql.FieldConfigArgument{"id": {Type: nonNull(gql.ID)}},
				Resolve: r.department,
			},
			"departments": {
				Type:    nonNull(gql.NewList(nonNull(departmentType))),
				Resolve: r.departmentList,
			},
		},
	})

	// ---------- Mutations (write) ----------
	mutation := gql.NewObject(gql.ObjectConfig{
		Name: "Mutation",
		Fields: gql.Fields{
			"createEmployee": {
				Type:    nonNull(employeeType),
				Args:    gql.FieldConfigArgument{"input": {Type: nonNull(createInput)}},
				Resolve: r.createEmployee,
			},
			"updateEmployee": {
				Type: nonNull(employeeType),
				Args: gql.FieldConfigArgument{
					"id":    {Type: nonNull(gql.ID)},
					"input": {Type: nonNull(updateInput)},
				},
				Resolve: r.updateEmployee,
			},
			"deleteEmployee": {
				Type:    nonNull(gql.Boolean),
				Args:    gql.FieldConfigArgument{"id": {Type: nonNull(gql.ID)}},
				Resolve: r.deleteEmployee,
			},
		},
	})

	return gql.NewSchema(gql.SchemaConfig{Query: query, Mutation: mutation})
}
