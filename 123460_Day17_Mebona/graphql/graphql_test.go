package graphql

import (
	"context"
	"employee-management/auth"
	"employee-management/graphql/resolver"
	"employee-management/graphql/schema"
	"encoding/json"
	gql "github.com/graphql-go/graphql"
	"strings"
	"testing"
)

// ---------- small helpers used by every test ----------

func newTestSchema(t *testing.T) (gql.Schema, *fakeEmployees) {
	t.Helper()
	emps := newFakeEmployees()
	r := resolver.New(emps, fakeDepartments{}, fakeAuth{}) // services are injected here
	s, err := schema.Build(r)
	if err != nil {
		t.Fatalf("schema could not be built: %v", err)
	}
	return s, emps
}

// ctxAs pretends that a user with the given role is logged in.
func ctxAs(role string) context.Context {
	return auth.WithUser(context.Background(), &auth.Claims{Username: "tester", Role: role})
}

func run(s gql.Schema, ctx context.Context, query string, vars map[string]interface{}) *gql.Result {
	return gql.Do(gql.Params{Schema: s, RequestString: query, VariableValues: vars, Context: ctx})
}

// errorCode returns extensions.code of the first error ("" when there is no error).
func errorCode(res *gql.Result) string {
	if len(res.Errors) == 0 {
		return ""
	}
	code, _ := res.Errors[0].Extensions["code"].(string)
	return code
}

func asJSON(t *testing.T, v interface{}) string {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("cannot convert result to JSON: %v", err)
	}
	return string(b)
}

func mustSucceed(t *testing.T, res *gql.Result) string {
	t.Helper()
	if len(res.Errors) > 0 {
		t.Fatalf("unexpected GraphQL errors: %v", res.Errors)
	}
	return asJSON(t, res.Data)
}

// ======================= QUERIES =======================

func TestQuery_Employee_WithNestedDepartment(t *testing.T) {
	s, _ := newTestSchema(t)
	res := run(s, ctxAs(auth.RoleUser), `{ employee(id: 1) { name email department { name } } }`, nil)
	out := mustSucceed(t, res)
	if !strings.Contains(out, "Amit") || !strings.Contains(out, "IT") {
		t.Fatalf("unexpected result: %s", out)
	}
}

func TestQuery_Employee_NotFound(t *testing.T) {
	s, _ := newTestSchema(t)
	res := run(s, ctxAs(auth.RoleUser), `{ employee(id: 999) { name } }`, nil)
	if code := errorCode(res); code != "NOT_FOUND" {
		t.Fatalf("expected NOT_FOUND, got %q", code)
	}
}

func TestQuery_Employee_InvalidID(t *testing.T) {
	s, _ := newTestSchema(t)
	res := run(s, ctxAs(auth.RoleUser), `{ employee(id: "abc") { name } }`, nil)
	if code := errorCode(res); code != "BAD_USER_INPUT" {
		t.Fatalf("expected BAD_USER_INPUT, got %q", code)
	}
}

func TestQuery_Employees_PaginationFilteringSorting(t *testing.T) {
	s, emps := newTestSchema(t)
	q := `{
	  employees(page: 2, pageSize: 5, search: "am", departmentId: 1, minSalary: 1000,
	            sortBy: SALARY, sortOrder: DESC) {
	    items { name }
	    totalCount page pageSize totalPages hasNext
	  }
	}`
	mustSucceed(t, run(s, ctxAs(auth.RoleUser), q, nil))

	// the resolver must have passed every argument to the service correctly
	f := emps.lastFilter
	if f.Page != 2 || f.PageSize != 5 {
		t.Errorf("paging wrong: page %d size %d", f.Page, f.PageSize)
	}
	if f.Search != "am" || f.DepartmentID != 1 {
		t.Errorf("filter wrong: search %q department %d", f.Search, f.DepartmentID)
	}
	if f.MinSalary == nil || *f.MinSalary != 1000 || f.MaxSalary != nil {
		t.Errorf("salary filter wrong: %v %v", f.MinSalary, f.MaxSalary)
	}
	if f.SortBy != "salary" || f.SortOrder != "desc" {
		t.Errorf("sorting wrong: %s %s", f.SortBy, f.SortOrder)
	}
}

func TestQuery_Employees_UsingVariables(t *testing.T) {
	s, emps := newTestSchema(t)
	q := `query List($page: Int, $text: String) { employees(page: $page, search: $text) { totalCount } }`
	mustSucceed(t, run(s, ctxAs(auth.RoleUser), q, map[string]interface{}{"page": 3, "text": "ri"}))
	if emps.lastFilter.Page != 3 || emps.lastFilter.Search != "ri" {
		t.Fatalf("variables were not used: %+v", emps.lastFilter)
	}
}

func TestQuery_DepartmentWithEmployees(t *testing.T) {
	s, _ := newTestSchema(t)
	out := mustSucceed(t, run(s, ctxAs(auth.RoleUser), `{ department(id: 1) { name employees { name } } }`, nil))
	if !strings.Contains(out, "Amit") {
		t.Fatalf("department should list Amit: %s", out)
	}
}

// ======================= MUTATIONS =======================

func TestMutation_CreateEmployee_AsAdmin(t *testing.T) {
	s, emps := newTestSchema(t)
	q := `mutation Create($input: EmployeeInput!) {
	  createEmployee(input: $input) { id name department { name } }
	}`
	vars := map[string]interface{}{"input": map[string]interface{}{
		"name": "Riya", "email": "riya@example.com", "salary": 82000.0, "departmentId": "1",
	}}
	out := mustSucceed(t, run(s, ctxAs(auth.RoleAdmin), q, vars))
	if !strings.Contains(out, "Riya") {
		t.Fatalf("unexpected result: %s", out)
	}
	if len(emps.items) != 2 {
		t.Fatalf("employee was not created")
	}
}

func TestMutation_UpdateEmployee(t *testing.T) {
	s, _ := newTestSchema(t)
	q := `mutation { updateEmployee(id: 1, input: {
	  name: "Amit Updated", email: "amit@example.com", salary: 80000, departmentId: 1
	}) { name salary } }`
	out := mustSucceed(t, run(s, ctxAs(auth.RoleAdmin), q, nil))
	if !strings.Contains(out, "Amit Updated") {
		t.Fatalf("unexpected result: %s", out)
	}
}

func TestMutation_UpdateEmployee_NotFound(t *testing.T) {
	s, _ := newTestSchema(t)
	q := `mutation { updateEmployee(id: 999, input: {
	  name: "Ghost", email: "g@example.com", salary: 1, departmentId: 1
	}) { id } }`
	if code := errorCode(run(s, ctxAs(auth.RoleAdmin), q, nil)); code != "NOT_FOUND" {
		t.Fatalf("expected NOT_FOUND, got %q", code)
	}
}

func TestMutation_DeleteEmployee(t *testing.T) {
	s, emps := newTestSchema(t)
	out := mustSucceed(t, run(s, ctxAs(auth.RoleAdmin), `mutation { deleteEmployee(id: 1) }`, nil))
	if !strings.Contains(out, "true") || len(emps.items) != 0 {
		t.Fatalf("employee was not deleted: %s", out)
	}
	// the employee must be gone now
	if code := errorCode(run(s, ctxAs(auth.RoleAdmin), `{ employee(id: 1) { name } }`, nil)); code != "NOT_FOUND" {
		t.Fatalf("expected NOT_FOUND after delete, got %q", code)
	}
}

func TestMutation_MissingRequiredFieldIsRejectedBySchema(t *testing.T) {
	s, _ := newTestSchema(t)
	q := `mutation { createEmployee(input: { name: "No email", salary: 1, departmentId: 1 }) { id } }`
	if res := run(s, ctxAs(auth.RoleAdmin), q, nil); len(res.Errors) == 0 {
		t.Fatal("a missing required field must produce an error")
	}
}

// ======================= SECURITY =======================

func TestSecurity_QueryNeedsLogin(t *testing.T) {
	s, _ := newTestSchema(t)
	res := run(s, context.Background(), `{ employees { totalCount } }`, nil)
	if code := errorCode(res); code != "UNAUTHENTICATED" {
		t.Fatalf("expected UNAUTHENTICATED, got %q", code)
	}
}

func TestSecurity_UserCannotCreateEmployee(t *testing.T) {
	s, emps := newTestSchema(t)
	q := `mutation { createEmployee(input: {
	  name: "Hacker", email: "h@example.com", salary: 1, departmentId: 1 }) { id } }`
	if code := errorCode(run(s, ctxAs(auth.RoleUser), q, nil)); code != "FORBIDDEN" {
		t.Fatalf("expected FORBIDDEN, got %q", code)
	}
	if len(emps.items) != 1 {
		t.Fatal("nothing must be created when access is forbidden")
	}
}

func TestSecurity_UserCannotDeleteEmployee(t *testing.T) {
	s, emps := newTestSchema(t)
	if code := errorCode(run(s, ctxAs(auth.RoleUser), `mutation { deleteEmployee(id: 1) }`, nil)); code != "FORBIDDEN" {
		t.Fatalf("expected FORBIDDEN, got %q", code)
	}
	if len(emps.items) != 1 {
		t.Fatal("employee must not be deleted by a normal user")
	}
}

func TestSecurity_LoginMutation(t *testing.T) {
	s, _ := newTestSchema(t)
	// login works WITHOUT a token
	out := mustSucceed(t, run(s, context.Background(), `mutation { login(username: "admin", password: "admin123") { token role } }`, nil))
	if !strings.Contains(out, "fake-token") || !strings.Contains(out, "ADMIN") {
		t.Fatalf("unexpected login result: %s", out)
	}
	// wrong password
	res := run(s, context.Background(), `mutation { login(username: "admin", password: "wrong") { token } }`, nil)
	if code := errorCode(res); code != "UNAUTHENTICATED" {
		t.Fatalf("expected UNAUTHENTICATED, got %q", code)
	}
}

func TestSecurity_Me(t *testing.T) {
	s, _ := newTestSchema(t)
	out := mustSucceed(t, run(s, ctxAs(auth.RoleUser), `{ me { username role } }`, nil))
	if !strings.Contains(out, "tester") || !strings.Contains(out, "USER") {
		t.Fatalf("unexpected result: %s", out)
	}
}
