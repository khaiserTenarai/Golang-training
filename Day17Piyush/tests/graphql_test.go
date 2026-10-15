// Tasks 12, 13: Test GraphQL queries and mutations

package tests

import (
	"context"
	"strings"
	"testing"

	"graphql_employee/schema"
	"graphql_employee/service"

	"github.com/graphql-go/graphql"
)

var testSchema graphql.Schema

func init() {
	empSvc := service.NewEmployeeService()
	deptSvc := service.NewDepartmentService()
	authSvc := service.NewAuthService()
	var err error
	testSchema, err = schema.BuildSchema(empSvc, deptSvc, authSvc)
	if err != nil {
		panic(err)
	}
}

func authCtx() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", 1)
	ctx = context.WithValue(ctx, "username", "admin")
	ctx = context.WithValue(ctx, "role", "admin")
	return ctx
}

func userCtx() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "user_id", 3)
	ctx = context.WithValue(ctx, "username", "piyush")
	ctx = context.WithValue(ctx, "role", "user")
	return ctx
}

// ---- Task 12: Test GraphQL Queries ----

func TestQueryEmployees(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ employees(page:1, page_size:3) { employees { id name } total } }`,
		Context:       authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("query error: %v", result.Errors)
	}
	data := result.Data.(map[string]interface{})
	empList := data["employees"].(map[string]interface{})
	total := empList["total"].(int)
	if total != 6 {
		t.Errorf("expected 6 employees, got %d", total)
	}
}

func TestQuerySingleEmployee(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ employee(id: 1) { id name email } }`,
		Context:       authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("query error: %v", result.Errors)
	}
	data := result.Data.(map[string]interface{})
	emp := data["employee"].(map[string]interface{})
	if emp["name"] != "Piyush" {
		t.Errorf("expected Piyush, got %v", emp["name"])
	}
}

func TestQueryEmployeeNotFound(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ employee(id: 999) { id name } }`,
		Context:       authCtx(),
	})
	if len(result.Errors) == 0 {
		t.Error("expected error for non-existent employee")
	}
}

func TestQueryNestedDepartment(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ employee(id: 1) { name department { name location } } }`,
		Context:       authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("query error: %v", result.Errors)
	}
	data := result.Data.(map[string]interface{})
	emp := data["employee"].(map[string]interface{})
	dept := emp["department"].(map[string]interface{})
	if dept["name"] != "Engineering" {
		t.Errorf("expected Engineering, got %v", dept["name"])
	}
}

func TestQueryWithPagination(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ employees(page:1, page_size:2) { employees { id } total total_pages } }`,
		Context:       authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("error: %v", result.Errors)
	}
	data := result.Data.(map[string]interface{})["employees"].(map[string]interface{})
	if data["total_pages"].(int) != 3 {
		t.Errorf("expected 3 pages, got %v", data["total_pages"])
	}
}

func TestQueryWithSorting(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ employees(sort_by:"salary", sort_order:"desc") { employees { name salary } } }`,
		Context:       authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("error: %v", result.Errors)
	}
}

func TestQueryRequiresAuth(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ employees { total } }`,
		Context:       context.Background(), // no auth
	})
	if len(result.Errors) == 0 {
		t.Error("expected auth error")
	}
}

func TestQueryDepartments(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `{ departments { id name location } }`,
		Context:       authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("error: %v", result.Errors)
	}
}

// ---- Task 13: Test GraphQL Mutations ----

func TestMutationLogin(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `mutation { login(username:"admin", password:"admin123") { token username role } }`,
	})
	if len(result.Errors) > 0 {
		t.Fatalf("login error: %v", result.Errors)
	}
	data := result.Data.(map[string]interface{})["login"].(map[string]interface{})
	if data["role"] != "admin" {
		t.Errorf("expected admin role, got %v", data["role"])
	}
	if data["token"] == nil || data["token"].(string) == "" {
		t.Error("expected non-empty token")
	}
}

func TestMutationLoginInvalid(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `mutation { login(username:"admin", password:"wrong") { token } }`,
	})
	if len(result.Errors) == 0 {
		t.Error("expected login error for wrong password")
	}
}

func TestMutationCreateEmployee(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema: testSchema,
		RequestString: `mutation {
			createEmployee(input: {name:"TestUser", email:"testuser@test.com", age:25, department_id:1, salary:50000}) {
				id name email
			}
		}`,
		Context: authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("create error: %v", result.Errors)
	}
	data := result.Data.(map[string]interface{})["createEmployee"].(map[string]interface{})
	if data["name"] != "TestUser" {
		t.Errorf("expected TestUser, got %v", data["name"])
	}
}

func TestMutationCreateDeniedForUser(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema: testSchema,
		RequestString: `mutation {
			createEmployee(input: {name:"Bad", email:"bad@test.com", age:25, department_id:1, salary:50000}) { id }
		}`,
		Context: userCtx(), // role=user, not admin/manager
	})
	if len(result.Errors) == 0 {
		t.Error("expected access denied for user role")
	}
	if !strings.Contains(result.Errors[0].Message, "access denied") {
		t.Errorf("expected 'access denied', got: %s", result.Errors[0].Message)
	}
}

func TestMutationCreateValidationFail(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema: testSchema,
		RequestString: `mutation {
			createEmployee(input: {name:"A", email:"bad", age:10, department_id:1, salary:-100}) { id }
		}`,
		Context: authCtx(),
	})
	if len(result.Errors) == 0 {
		t.Error("expected validation error")
	}
}

func TestMutationDeleteEmployee(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `mutation { deleteEmployee(id: 6) }`,
		Context:       authCtx(),
	})
	if len(result.Errors) > 0 {
		t.Fatalf("delete error: %v", result.Errors)
	}
}

func TestMutationDeleteDeniedForUser(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `mutation { deleteEmployee(id: 1) }`,
		Context:       userCtx(),
	})
	if len(result.Errors) == 0 {
		t.Error("expected access denied")
	}
}

func TestMutationRegister(t *testing.T) {
	result := graphql.Do(graphql.Params{
		Schema:        testSchema,
		RequestString: `mutation { register(username:"newuser", password:"newpass123") { token username role } }`,
	})
	if len(result.Errors) > 0 {
		t.Fatalf("register error: %v", result.Errors)
	}
}
