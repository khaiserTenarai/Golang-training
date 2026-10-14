package schema_test

import (
	"context"
	"testing"

	"employee-platform/models"
	"employee-platform/repository"
	"employee-platform/schema"
	"employee-platform/service"

	"github.com/graphql-go/graphql"
)

func TestGraphQL_EmployeesQuery(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewEmployeeService(repo)
	gqlSchema, err := schema.BuildSchema(svc)
	if err != nil {
		t.Fatalf("failed to build schema: %v", err)
	}

	query := `
		query {
			employees(limit: 2) {
				id
				name
				department {
					name
				}
			}
		}
	`

	result := graphql.Do(graphql.Params{
		Schema:        gqlSchema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		t.Fatalf("graphql query errors: %v", result.Errors)
	}
}

func TestGraphQL_CreateEmployeeMutation(t *testing.T) {
	repo := repository.NewMemoryRepository()
	svc := service.NewEmployeeService(repo)
	gqlSchema, _ := schema.BuildSchema(svc)

	mutation := `
		mutation {
			createEmployee(id: "emp-99", name: "New Hire", role: "Dev", departmentId: "dept-1", salary: 90000) {
				id
				name
			}
		}
	`

	// Context with ADMIN role
	ctx := context.WithValue(context.Background(), "userCtx", models.UserContext{UserID: "usr-admin", Role: "ADMIN"})

	result := graphql.Do(graphql.Params{
		Schema:        gqlSchema,
		RequestString: mutation,
		Context:       ctx,
	})

	if len(result.Errors) > 0 {
		t.Fatalf("mutation errors: %v", result.Errors)
	}
}
