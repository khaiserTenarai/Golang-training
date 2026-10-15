package graphql

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"employee-management/service"

	gql "github.com/graphql-go/graphql"
)

func newServer(t *testing.T) *Server {
	t.Helper()
	depts := newFakeDeptRepo()
	emps := newFakeEmpRepo(depts)
	s, err := NewServer(service.NewEmployeeService(emps, depts), service.NewDepartmentService(depts, emps))
	if err != nil {
		t.Fatal(err)
	}
	return s
}

// run executes a query and returns the whole response as a generic map
// ({"data": ..., "errors": [...]}) so tests can inspect it like a client would.
func run(t *testing.T, s *Server, query string, vars map[string]interface{}) map[string]interface{} {
	t.Helper()
	res := gql.Do(gql.Params{Schema: s.schema, RequestString: query, VariableValues: vars})
	raw, _ := json.Marshal(res)
	var out map[string]interface{}
	_ = json.Unmarshal(raw, &out)
	return out
}

// dig walks nested maps/slices: dig(m, "data", "employees", 0, "name")
func dig(v interface{}, path ...interface{}) interface{} {
	for _, p := range path {
		switch key := p.(type) {
		case string:
			m, ok := v.(map[string]interface{})
			if !ok {
				return nil
			}
			v = m[key]
		case int:
			a, ok := v.([]interface{})
			if !ok || key >= len(a) {
				return nil
			}
			v = a[key]
		}
	}
	return v
}

func errCode(res map[string]interface{}) interface{} {
	return dig(res, "errors", 0, "extensions", "code")
}

func TestEmployeeList(t *testing.T) {
	s := newServer(t)
	res := run(t, s, `{ employees { id name department { name } } }`, nil)
	if res["errors"] != nil {
		t.Fatalf("unexpected errors: %v", res["errors"])
	}
	if got := len(dig(res, "data", "employees").([]interface{})); got != 3 {
		t.Fatalf("want 3 employees, got %d", got)
	}
	if dig(res, "data", "employees", 0, "department", "name") != "Engineering" {
		t.Fatalf("nested department not resolved: %v", res)
	}
}

func TestEmployeeByIDWithVariable(t *testing.T) {
	s := newServer(t)
	q := `query GetEmployee($id: ID!) { employee(id: $id) { name email salary joinedAt } }`
	res := run(t, s, q, map[string]interface{}{"id": "2"})
	if dig(res, "data", "employee", "name") != "Amit" {
		t.Fatalf("unexpected: %v", res)
	}
	if dig(res, "data", "employee", "joinedAt") != "2023-03-15T00:00:00Z" {
		t.Fatalf("DateTime scalar not serialised: %v", res)
	}
}

func TestEmployeeNotFound(t *testing.T) {
	s := newServer(t)
	res := run(t, s, `{ employee(id: "99") { id name } }`, nil)
	if dig(res, "data", "employee") != nil {
		t.Fatalf("employee should be null: %v", res)
	}
	if errCode(res) != "NOT_FOUND" {
		t.Fatalf("want NOT_FOUND, got %v (%v)", errCode(res), res["errors"])
	}
	res = run(t, s, `{ employee(id: "abc") { id } }`, nil)
	if errCode(res) != "BAD_USER_INPUT" {
		t.Fatalf("non-numeric id: want BAD_USER_INPUT, got %v", res["errors"])
	}
}

func TestListArguments(t *testing.T) {
	s := newServer(t)

	res := run(t, s, `{ employees(departmentId: "1") { name } }`, nil)
	if got := len(dig(res, "data", "employees").([]interface{})); got != 2 {
		t.Fatalf("department filter: want 2, got %d", got)
	}

	res = run(t, s, `{ employees(limit: 1, offset: 1) { name } }`, nil)
	if dig(res, "data", "employees", 0, "name") != "Amit" {
		t.Fatalf("pagination failed: %v", res)
	}

	res = run(t, s, `query($after: DateTime) { employees(joinedAfter: $after) { name } }`,
		map[string]interface{}{"after": "2023-01-01T00:00:00Z"})
	if got := len(dig(res, "data", "employees").([]interface{})); got != 2 {
		t.Fatalf("joinedAfter filter: want 2, got %d (%v)", got, res)
	}

	res = run(t, s, `{ employees(search: "pri") { name } }`, nil)
	if dig(res, "data", "employees", 0, "name") != "Priya" {
		t.Fatalf("search failed: %v", res)
	}

	res = run(t, s, `{ employees(limit: 1000) { name } }`, nil)
	if errCode(res) != "BAD_USER_INPUT" {
		t.Fatalf("want BAD_USER_INPUT for limit=1000, got %v", res)
	}
}

const createMutation = `
mutation Create($input: CreateEmployeeInput!) {
  createEmployee(input: $input) { id name email salary active department { name } }
}`

func input(overrides map[string]interface{}) map[string]interface{} {
	in := map[string]interface{}{
		"name": "John", "email": "john@gmail.com", "salary": 70000, "departmentId": "3",
	}
	for k, v := range overrides {
		in[k] = v
	}
	return map[string]interface{}{"input": in}
}

func TestCreateEmployee(t *testing.T) {
	s := newServer(t)
	res := run(t, s, createMutation, input(nil))
	if res["errors"] != nil {
		t.Fatalf("unexpected errors: %v", res["errors"])
	}
	if dig(res, "data", "createEmployee", "id") != "4" ||
		dig(res, "data", "createEmployee", "active") != true || // input default value
		dig(res, "data", "createEmployee", "department", "name") != "Finance" {
		t.Fatalf("unexpected create result: %v", res)
	}
}

func TestCreateValidationAndConflict(t *testing.T) {
	s := newServer(t)
	cases := []struct {
		name, code, field string
		over              map[string]interface{}
	}{
		{"bad email", "BAD_USER_INPUT", "email", map[string]interface{}{"email": "nope"}},
		{"short name", "BAD_USER_INPUT", "name", map[string]interface{}{"name": "J"}},
		{"negative salary", "BAD_USER_INPUT", "salary", map[string]interface{}{"salary": -1}},
		{"unknown department", "BAD_USER_INPUT", "departmentId", map[string]interface{}{"departmentId": "42"}},
		{"duplicate email", "CONFLICT", "email", map[string]interface{}{"email": "rajesh@gmail.com"}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res := run(t, s, createMutation, input(tc.over))
			if errCode(res) != tc.code || dig(res, "errors", 0, "extensions", "field") != tc.field {
				t.Fatalf("want %s on %s, got %v", tc.code, tc.field, res["errors"])
			}
		})
	}
}

// Missing required input fields / unknown fields are rejected by GraphQL itself,
// before any resolver runs.
func TestSchemaLevelValidation(t *testing.T) {
	s := newServer(t)
	res := run(t, s, createMutation, map[string]interface{}{
		"input": map[string]interface{}{"name": "John"},
	})
	if res["errors"] == nil {
		t.Fatal("expected a GraphQL validation error")
	}
	if errCode(res) != nil {
		t.Fatal("schema-level errors carry no custom code")
	}
	if res = run(t, s, `{ employees { doesNotExist } }`, nil); res["errors"] == nil {
		t.Fatal("unknown field must be rejected")
	}
}

func TestUpdateEmployee(t *testing.T) {
	s := newServer(t)
	q := `mutation($id: ID!, $input: UpdateEmployeeInput!) {
	  updateEmployee(id: $id, input: $input) { name email salary department { name } }
	}`

	// partial update: only salary and department are sent
	res := run(t, s, q, map[string]interface{}{"id": "1",
		"input": map[string]interface{}{"salary": 80000, "departmentId": "3"}})
	if dig(res, "data", "updateEmployee", "salary") != float64(80000) ||
		dig(res, "data", "updateEmployee", "name") != "Rajesh" ||
		dig(res, "data", "updateEmployee", "department", "name") != "Finance" {
		t.Fatalf("partial update failed: %v", res)
	}

	res = run(t, s, q, map[string]interface{}{"id": "99", "input": map[string]interface{}{"salary": 1}})
	if errCode(res) != "NOT_FOUND" {
		t.Fatalf("want NOT_FOUND, got %v", res["errors"])
	}

	res = run(t, s, q, map[string]interface{}{"id": "1", "input": map[string]interface{}{"email": "amit@gmail.com"}})
	if errCode(res) != "CONFLICT" {
		t.Fatalf("want CONFLICT, got %v", res["errors"])
	}

	// keeping your own email is not a conflict
	res = run(t, s, q, map[string]interface{}{"id": "1", "input": map[string]interface{}{"email": "rajesh@gmail.com"}})
	if res["errors"] != nil {
		t.Fatalf("own email must be allowed: %v", res["errors"])
	}
}

func TestDeleteEmployee(t *testing.T) {
	s := newServer(t)
	q := `mutation($id: ID!) { deleteEmployee(id: $id) }`

	if res := run(t, s, q, map[string]interface{}{"id": "2"}); dig(res, "data", "deleteEmployee") != true {
		t.Fatalf("delete failed: %v", res)
	}
	if res := run(t, s, `{ employee(id: "2") { id } }`, nil); errCode(res) != "NOT_FOUND" {
		t.Fatalf("deleted employee should be gone: %v", res)
	}
	if res := run(t, s, q, map[string]interface{}{"id": "2"}); errCode(res) != "NOT_FOUND" {
		t.Fatalf("second delete should be NOT_FOUND: %v", res)
	}
}

func TestDepartments(t *testing.T) {
	s := newServer(t)
	res := run(t, s, `{ department(id: "1") { name employeeCount employees { name } } }`, nil)
	if dig(res, "data", "department", "employeeCount") != float64(2) {
		t.Fatalf("unexpected: %v", res)
	}
	res = run(t, s, `{ departments { name } }`, nil)
	if got := len(dig(res, "data", "departments").([]interface{})); got != 3 {
		t.Fatalf("want 3 departments, got %d", got)
	}
}

// End-to-end through the real HTTP handler, the way a client would call it.
func TestHTTPEndpoint(t *testing.T) {
	h := NewHTTPHandler(newServer(t))
	body, _ := json.Marshal(map[string]interface{}{
		"query":     `query($id: ID!) { employee(id: $id) { name } }`,
		"variables": map[string]interface{}{"id": "1"},
	})
	req := httptest.NewRequest(http.MethodPost, "/graphql", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	var res map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &res)
	if w.Code != http.StatusOK || dig(res, "data", "employee", "name") != "Rajesh" {
		t.Fatalf("unexpected response: %d %s", w.Code, w.Body.String())
	}
}
