package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

type fakeRepo struct {
	mu         sync.Mutex
	employees  map[int]Employee
	nextID     int
	lastParams ListParams
	pingErr    error
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{employees: make(map[int]Employee), nextID: 1}
}

func (f *fakeRepo) emailTaken(email string, exceptID int) bool {
	for _, e := range f.employees {
		if e.ID != exceptID && e.Email == email {
			return true
		}
	}
	return false
}

func (f *fakeRepo) Create(ctx context.Context, in EmployeeInput) (Employee, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.emailTaken(in.Email, 0) {
		return Employee{}, ErrDuplicateEmail
	}
	now := time.Now()
	e := Employee{ID: f.nextID, Name: in.Name, Email: in.Email, Department: in.Department, Salary: in.Salary, CreatedAt: now, UpdatedAt: now}
	f.employees[e.ID] = e
	f.nextID++
	return e, nil
}

func (f *fakeRepo) GetByID(ctx context.Context, id int) (Employee, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	e, ok := f.employees[id]
	if !ok {
		return Employee{}, ErrNotFound
	}
	return e, nil
}

func (f *fakeRepo) List(ctx context.Context, p ListParams) ([]Employee, int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.lastParams = p

	var matched []Employee
	for id := 1; id < f.nextID; id++ {
		e, ok := f.employees[id]
		if !ok {
			continue
		}
		if p.Department != "" && !strings.EqualFold(e.Department, p.Department) {
			continue
		}
		matched = append(matched, e)
	}

	total := len(matched)
	start := (p.Page - 1) * p.Limit
	if start > total {
		start = total
	}
	end := start + p.Limit
	if end > total {
		end = total
	}
	result := matched[start:end]
	if result == nil {
		result = []Employee{}
	}
	return result, total, nil
}

func (f *fakeRepo) Update(ctx context.Context, id int, in EmployeeInput) (Employee, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	e, ok := f.employees[id]
	if !ok {
		return Employee{}, ErrNotFound
	}
	if f.emailTaken(in.Email, id) {
		return Employee{}, ErrDuplicateEmail
	}
	e.Name, e.Email, e.Department, e.Salary, e.UpdatedAt = in.Name, in.Email, in.Department, in.Salary, time.Now()
	f.employees[id] = e
	return e, nil
}

func (f *fakeRepo) Delete(ctx context.Context, id int) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if _, ok := f.employees[id]; !ok {
		return ErrNotFound
	}
	delete(f.employees, id)
	return nil
}

func (f *fakeRepo) Ping(ctx context.Context) error {
	return f.pingErr
}

func do(t *testing.T, h http.Handler, method, path, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/json")
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func expectStatus(t *testing.T, rec *httptest.ResponseRecorder, want int) {
	t.Helper()
	if rec.Code != want {
		t.Fatalf("status = %d, want %d, body = %s", rec.Code, want, rec.Body.String())
	}
}

const validBody = `{"name":"Anita Rao","email":"anita@example.com","department":"Engineering","salary":85000}`

func TestCreateEmployee(t *testing.T) {
	h := NewRouter(newFakeRepo())

	rec := do(t, h, http.MethodPost, "/employee", validBody)
	expectStatus(t, rec, http.StatusCreated)

	if got := rec.Header().Get("Location"); got != "/employee/1" {
		t.Errorf("Location = %q, want /employee/1", got)
	}
	if got := rec.Header().Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q", got)
	}

	var e Employee
	if err := json.Unmarshal(rec.Body.Bytes(), &e); err != nil {
		t.Fatal(err)
	}
	if e.ID != 1 || e.Name != "Anita Rao" || e.Salary != 85000 {
		t.Errorf("unexpected employee: %+v", e)
	}
}

func TestCreateNormalizesInput(t *testing.T) {
	h := NewRouter(newFakeRepo())

	rec := do(t, h, http.MethodPost, "/employee", `{"name":"  Anita Rao ","email":" ANITA@Example.com ","department":"Engineering","salary":1}`)
	expectStatus(t, rec, http.StatusCreated)

	var e Employee
	json.Unmarshal(rec.Body.Bytes(), &e)
	if e.Name != "Anita Rao" || e.Email != "anita@example.com" {
		t.Errorf("input not normalized: %+v", e)
	}
}

func TestCreateValidationErrors(t *testing.T) {
	h := NewRouter(newFakeRepo())

	rec := do(t, h, http.MethodPost, "/employee", `{"name":"A","email":"bad","department":"Pirates","salary":-1}`)
	expectStatus(t, rec, http.StatusUnprocessableEntity)

	var body errorBody
	json.Unmarshal(rec.Body.Bytes(), &body)
	for _, field := range []string{"name", "email", "department", "salary"} {
		if body.Fields[field] == "" {
			t.Errorf("expected validation error for %q, got %v", field, body.Fields)
		}
	}
}

func TestCreateBadRequests(t *testing.T) {
	h := NewRouter(newFakeRepo())

	cases := []struct {
		name string
		body string
		want int
	}{
		{"malformed json", `{bad`, http.StatusBadRequest},
		{"empty body", ``, http.StatusBadRequest},
		{"unknown field", `{"name":"Anita Rao","email":"a@b.com","department":"HR","salary":1,"id":99}`, http.StatusBadRequest},
		{"wrong type", `{"name":"Anita Rao","email":"a@b.com","department":"HR","salary":"lots"}`, http.StatusBadRequest},
		{"trailing data", validBody + `{"x":1}`, http.StatusBadRequest},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			expectStatus(t, do(t, h, http.MethodPost, "/employee", c.body), c.want)
		})
	}
}

func TestCreateRequiresJSONContentType(t *testing.T) {
	h := NewRouter(newFakeRepo())

	req := httptest.NewRequest(http.MethodPost, "/employee", strings.NewReader(validBody))
	req.Header.Set("Content-Type", "text/plain")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)

	expectStatus(t, rec, http.StatusUnsupportedMediaType)
}

func TestCreateDuplicateEmail(t *testing.T) {
	h := NewRouter(newFakeRepo())

	expectStatus(t, do(t, h, http.MethodPost, "/employee", validBody), http.StatusCreated)
	expectStatus(t, do(t, h, http.MethodPost, "/employee", validBody), http.StatusConflict)
}

func TestGetEmployee(t *testing.T) {
	h := NewRouter(newFakeRepo())
	do(t, h, http.MethodPost, "/employee", validBody)

	expectStatus(t, do(t, h, http.MethodGet, "/employee/1", ""), http.StatusOK)
	expectStatus(t, do(t, h, http.MethodGet, "/employee/99", ""), http.StatusNotFound)
	expectStatus(t, do(t, h, http.MethodGet, "/employee/abc", ""), http.StatusBadRequest)
	expectStatus(t, do(t, h, http.MethodGet, "/employee/0", ""), http.StatusBadRequest)
}

func TestUpdateEmployee(t *testing.T) {
	h := NewRouter(newFakeRepo())
	do(t, h, http.MethodPost, "/employee", validBody)

	rec := do(t, h, http.MethodPut, "/employee/1", `{"name":"Anita R","email":"anita@example.com","department":"Finance","salary":90000}`)
	expectStatus(t, rec, http.StatusOK)

	var e Employee
	json.Unmarshal(rec.Body.Bytes(), &e)
	if e.Name != "Anita R" || e.Department != "Finance" || e.Salary != 90000 {
		t.Errorf("update not applied: %+v", e)
	}

	expectStatus(t, do(t, h, http.MethodPut, "/employee/99", validBody), http.StatusNotFound)
	expectStatus(t, do(t, h, http.MethodPut, "/employee/1", `{"name":"x"}`), http.StatusUnprocessableEntity)
}

func TestUpdateDuplicateEmail(t *testing.T) {
	h := NewRouter(newFakeRepo())
	do(t, h, http.MethodPost, "/employee", validBody)
	do(t, h, http.MethodPost, "/employee", `{"name":"Rahul Verma","email":"rahul@example.com","department":"Finance","salary":1}`)

	rec := do(t, h, http.MethodPut, "/employee/2", `{"name":"Rahul Verma","email":"anita@example.com","department":"Finance","salary":1}`)
	expectStatus(t, rec, http.StatusConflict)
}

func TestDeleteEmployee(t *testing.T) {
	h := NewRouter(newFakeRepo())
	do(t, h, http.MethodPost, "/employee", validBody)

	rec := do(t, h, http.MethodDelete, "/employee/1", "")
	expectStatus(t, rec, http.StatusNoContent)
	if rec.Body.Len() != 0 {
		t.Errorf("204 response must have empty body, got %q", rec.Body.String())
	}

	expectStatus(t, do(t, h, http.MethodDelete, "/employee/1", ""), http.StatusNotFound)
	expectStatus(t, do(t, h, http.MethodGet, "/employee/1", ""), http.StatusNotFound)
}

func TestListPaginationAndFilters(t *testing.T) {
	repo := newFakeRepo()
	h := NewRouter(repo)

	for i, dept := range []string{"Engineering", "Finance", "Engineering", "Sales", "Engineering"} {
		body := `{"name":"Person ` + string(rune('A'+i)) + `","email":"p` + string(rune('a'+i)) + `@example.com","department":"` + dept + `","salary":1000}`
		expectStatus(t, do(t, h, http.MethodPost, "/employee", body), http.StatusCreated)
	}

	rec := do(t, h, http.MethodGet, "/employee?page=2&limit=2", "")
	expectStatus(t, rec, http.StatusOK)

	var resp listResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	if len(resp.Data) != 2 || resp.Meta.Total != 5 || resp.Meta.TotalPages != 3 || resp.Meta.Page != 2 {
		t.Errorf("unexpected pagination: %+v", resp.Meta)
	}

	rec = do(t, h, http.MethodGet, "/employee?department=engineering&search=bob&sort=salary&order=DESC", "")
	expectStatus(t, rec, http.StatusOK)
	want := ListParams{Page: 1, Limit: 10, Department: "engineering", Search: "bob", Sort: "salary", Order: "desc"}
	if repo.lastParams != want {
		t.Errorf("params = %+v, want %+v", repo.lastParams, want)
	}

	rec = do(t, h, http.MethodGet, "/employee?department=Nowhere", "")
	expectStatus(t, rec, http.StatusOK)
	if !strings.Contains(rec.Body.String(), `"data":[]`) {
		t.Errorf("empty result must be [] not null: %s", rec.Body.String())
	}
}

func TestListInvalidParams(t *testing.T) {
	h := NewRouter(newFakeRepo())

	for _, query := range []string{"page=0", "page=abc", "limit=0", "limit=101", "sort=bogus", "order=sideways"} {
		t.Run(query, func(t *testing.T) {
			expectStatus(t, do(t, h, http.MethodGet, "/employee?"+query, ""), http.StatusBadRequest)
		})
	}
}

func TestMethodNotAllowed(t *testing.T) {
	h := NewRouter(newFakeRepo())

	rec := do(t, h, http.MethodDelete, "/employee", "")
	expectStatus(t, rec, http.StatusMethodNotAllowed)
	if rec.Header().Get("Allow") != "GET, POST" {
		t.Errorf("Allow = %q", rec.Header().Get("Allow"))
	}

	rec = do(t, h, http.MethodPatch, "/employee/1", "")
	expectStatus(t, rec, http.StatusMethodNotAllowed)
	if rec.Header().Get("Allow") != "GET, PUT, DELETE" {
		t.Errorf("Allow = %q", rec.Header().Get("Allow"))
	}
}

func TestUnknownRoute(t *testing.T) {
	h := NewRouter(newFakeRepo())
	rec := do(t, h, http.MethodGet, "/nope", "")
	expectStatus(t, rec, http.StatusNotFound)
	if !strings.Contains(rec.Body.String(), "route not found") {
		t.Errorf("body = %s", rec.Body.String())
	}
}

func TestHealth(t *testing.T) {
	repo := newFakeRepo()
	h := NewRouter(repo)

	expectStatus(t, do(t, h, http.MethodGet, "/health", ""), http.StatusOK)

	repo.pingErr = context.DeadlineExceeded
	expectStatus(t, do(t, h, http.MethodGet, "/health", ""), http.StatusServiceUnavailable)
}

func TestRequestIDHeader(t *testing.T) {
	h := NewRouter(newFakeRepo())
	rec := do(t, h, http.MethodGet, "/health", "")
	if !strings.HasPrefix(rec.Header().Get("X-Request-ID"), "req-") {
		t.Errorf("missing X-Request-ID header")
	}
}

func TestPanicRecovery(t *testing.T) {
	boom := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { panic("boom") })
	h := loggingMiddleware(recoveryMiddleware(boom))

	rec := do(t, h, http.MethodGet, "/", "")
	expectStatus(t, rec, http.StatusInternalServerError)
}

func TestPostgresRepositoryIntegration(t *testing.T) {
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}

	ctx := context.Background()
	db, err := openDB(ctx, url)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	if _, err := db.ExecContext(ctx, `DROP TABLE IF EXISTS employees`); err != nil {
		t.Fatal(err)
	}
	if err := Migrate(ctx, db); err != nil {
		t.Fatal(err)
	}

	repo := NewPostgresRepository(db)

	a, err := repo.Create(ctx, EmployeeInput{Name: "Anita Rao", Email: "anita@example.com", Department: "Engineering", Salary: 85000.5})
	if err != nil {
		t.Fatal(err)
	}
	if a.ID == 0 || a.CreatedAt.IsZero() {
		t.Fatalf("create did not populate fields: %+v", a)
	}

	if _, err := repo.Create(ctx, EmployeeInput{Name: "Other", Email: "anita@example.com", Department: "HR", Salary: 1}); err != ErrDuplicateEmail {
		t.Fatalf("duplicate email error = %v, want ErrDuplicateEmail", err)
	}

	b, _ := repo.Create(ctx, EmployeeInput{Name: "Rahul Verma", Email: "rahul@example.com", Department: "Finance", Salary: 62000})
	repo.Create(ctx, EmployeeInput{Name: "100% Real_Name", Email: "real@example.com", Department: "Engineering", Salary: 70000})

	got, err := repo.GetByID(ctx, a.ID)
	if err != nil || got.Salary != 85000.5 {
		t.Fatalf("get = %+v, %v", got, err)
	}
	if _, err := repo.GetByID(ctx, 9999); err != ErrNotFound {
		t.Fatalf("get missing error = %v", err)
	}

	updated, err := repo.Update(ctx, b.ID, EmployeeInput{Name: "Rahul V", Email: "rahul@example.com", Department: "Sales", Salary: 65000})
	if err != nil || updated.Name != "Rahul V" || !updated.UpdatedAt.After(b.UpdatedAt) {
		t.Fatalf("update = %+v, %v", updated, err)
	}
	if _, err := repo.Update(ctx, 9999, EmployeeInput{Name: "x", Email: "x@x.com", Department: "HR", Salary: 1}); err != ErrNotFound {
		t.Fatalf("update missing error = %v", err)
	}
	if _, err := repo.Update(ctx, b.ID, EmployeeInput{Name: "Rahul V", Email: "anita@example.com", Department: "Sales", Salary: 1}); err != ErrDuplicateEmail {
		t.Fatalf("update duplicate error = %v", err)
	}

	list, total, err := repo.List(ctx, ListParams{Page: 1, Limit: 10, Department: "engineering", Sort: "salary", Order: "desc"})
	if err != nil || total != 2 || len(list) != 2 || list[0].Salary < list[1].Salary {
		t.Fatalf("filtered list = %+v total=%d err=%v", list, total, err)
	}

	list, total, _ = repo.List(ctx, ListParams{Page: 1, Limit: 10, Search: "100%"})
	if total != 1 || list[0].Email != "real@example.com" {
		t.Fatalf("LIKE wildcards must be escaped: %+v total=%d", list, total)
	}

	list, total, _ = repo.List(ctx, ListParams{Page: 2, Limit: 2, Sort: "id", Order: "asc"})
	if total != 3 || len(list) != 1 {
		t.Fatalf("page 2 = %+v total=%d", list, total)
	}

	if err := repo.Delete(ctx, a.ID); err != nil {
		t.Fatal(err)
	}
	if err := repo.Delete(ctx, a.ID); err != ErrNotFound {
		t.Fatalf("delete twice error = %v", err)
	}

	if err := repo.Ping(ctx); err != nil {
		t.Fatal(err)
	}

	var count int
	if err := db.QueryRowContext(ctx, `SELECT COUNT(*) FROM employees`).Scan(&count); err != nil && err != sql.ErrNoRows {
		t.Fatal(err)
	}
	if count != 2 {
		t.Fatalf("row count = %d, want 2", count)
	}
}
