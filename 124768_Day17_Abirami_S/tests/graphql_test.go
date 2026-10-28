package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"day17-employee-platform/config"
	"day17-employee-platform/database"
	"day17-employee-platform/graphql"
	"day17-employee-platform/repository"
	"day17-employee-platform/service"

	gql "github.com/graphql-go/graphql"
)

func newGraphQLServerForTest(t *testing.T) *graphql.Server {
	t.Helper()

	cfg := config.Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "Root",
		DBName:     "day17_employee_db",
	}

	db, err := database.Connect(cfg)
	if err != nil {
		t.Skip("PostgreSQL is not available")
	}

	if err := database.Migrate(db); err != nil {
		t.Skip("PostgreSQL database is not available")
	}

	employeeRepo := repository.NewEmployeeRepository(db)
	departmentRepo := repository.NewDepartmentRepository(db)
	employeeService := service.NewEmployeeService(employeeRepo)
	departmentService := service.NewDepartmentService(departmentRepo)

	server, err := graphql.NewServer(employeeService, departmentService)
	if err != nil {
		t.Fatal(err)
	}

	return server
}

func TestGraphQLQuery(t *testing.T) {
	server := newGraphQLServerForTest(t)

	result := gql.Do(gql.Params{
		Schema: server.Schema,
		RequestString: `query {
			employees(page: 1, pageSize: 5) {
				items {
					id
					name
					department { id name }
				}
				total
			}
		}`,
	})

	if len(result.Errors) > 0 {
		t.Fatal(result.Errors)
	}
}

func TestGraphQLMutation(t *testing.T) {
	server := newGraphQLServerForTest(t)

	email := fmt.Sprintf("graphql-test-%d@example.com", time.Now().UnixNano())

	result := gql.Do(gql.Params{
		Schema:  server.Schema,
		Context: graphql.WithRoleContext(context.Background(), "admin"),
		RequestString: fmt.Sprintf(`mutation {
			createEmployee(
				input: {
					name: "GraphQL Test"
					email: "%s"
					salary: 30000
					departmentId: 1
				}
			) {
				id
				name
				email
			}
		}`, email),
	})

	if len(result.Errors) > 0 {
		t.Fatal(result.Errors)
	}
}
