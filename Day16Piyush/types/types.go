// Tasks 2, 3, 9, 10: GraphQL Type Definitions
// Task 2  — Create Employee type
// Task 3  — Create Department type
// Task 9  — Use GraphQL scalar types (Int, Float, String, Boolean, ID)
// Task 10 — Create input types (for mutations)

package types

import "github.com/graphql-go/graphql"

// Task 3: Department Type
// Task 9: Scalar types — Int, String, Float
var DepartmentType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Department",
	Description: "A company department",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.Int, Description: "Unique department ID"},                // Task 9: Int scalar
		"name":     &graphql.Field{Type: graphql.String, Description: "Department name"},                  // Task 9: String scalar
		"location": &graphql.Field{Type: graphql.String, Description: "Office location"},
		"budget":   &graphql.Field{Type: graphql.Float, Description: "Annual budget"},                     // Task 9: Float scalar
	},
})

// Task 2: Employee Type
// Task 9: Uses all scalar types — Int, Float, String, Boolean, ID
var EmployeeType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Employee",
	Description: "An employee in the system",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.NewNonNull(graphql.Int), Description: "Unique ID"},   // Task 9: NonNull Int
		"name":          &graphql.Field{Type: graphql.String, Description: "Full name"},
		"email":         &graphql.Field{Type: graphql.String, Description: "Email address"},
		"age":           &graphql.Field{Type: graphql.Int, Description: "Age in years"},
		"department_id": &graphql.Field{Type: graphql.Int, Description: "FK to department"},
		"salary":        &graphql.Field{Type: graphql.Float, Description: "Monthly salary"},               // Task 9: Float
		"is_active":     &graphql.Field{Type: graphql.Boolean, Description: "Employment status"},          // Task 9: Boolean scalar
		"join_date":     &graphql.Field{Type: graphql.String, Description: "Date joined (YYYY-MM-DD)"},
		"created_at":    &graphql.Field{Type: graphql.String, Description: "Record creation timestamp"},
	},
})

// Task 10: Input type for Create Employee mutation
var CreateEmployeeInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "CreateEmployeeInput",
	Description: "Input for creating a new employee",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "Required: full name"},
		"email":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String), Description: "Required: email"},
		"age":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int), Description: "Required: age (18-65)"},
		"department_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int), Description: "Required: department ID"},
		"salary":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float), Description: "Required: salary"},
		"is_active":     &graphql.InputObjectFieldConfig{Type: graphql.Boolean, DefaultValue: true},
		"join_date":     &graphql.InputObjectFieldConfig{Type: graphql.String, Description: "YYYY-MM-DD (defaults to today)"},
	},
})

// Task 10: Input type for Update Employee mutation
var UpdateEmployeeInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "UpdateEmployeeInput",
	Description: "Input for updating an employee",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":          &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"email":         &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
		"age":           &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"department_id": &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Int)},
		"salary":        &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.Float)},
		"is_active":     &graphql.InputObjectFieldConfig{Type: graphql.Boolean},
		"join_date":     &graphql.InputObjectFieldConfig{Type: graphql.String},
	},
})

// Employee list result type (for list query)
var EmployeeListType = graphql.NewObject(graphql.ObjectConfig{
	Name: "EmployeeList",
	Fields: graphql.Fields{
		"employees": &graphql.Field{Type: graphql.NewList(EmployeeType)},
		"count":     &graphql.Field{Type: graphql.Int},
	},
})

// Delete result type
var DeleteResultType = graphql.NewObject(graphql.ObjectConfig{
	Name: "DeleteResult",
	Fields: graphql.Fields{
		"success": &graphql.Field{Type: graphql.Boolean},
		"message": &graphql.Field{Type: graphql.String},
	},
})
