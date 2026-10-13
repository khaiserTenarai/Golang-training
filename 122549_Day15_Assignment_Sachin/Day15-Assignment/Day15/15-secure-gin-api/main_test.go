package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)


func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func newTestRouter(t *testing.T) *gin.Engine {
	auth := NewAuth("test-secret")
	if err := auth.AddUser("alice", "adminpass", "admin"); err != nil {
		t.Fatal(err)
	}
	if err := auth.AddUser("bob", "userpass", "user"); err != nil {
		t.Fatal(err)
	}
	return NewRouter(auth, NewStore())
}

func send(router http.Handler, method, path, token, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func loginAs(t *testing.T, router http.Handler, username, password string) string {
	body := fmt.Sprintf(`{"username":%q,"password":%q}`, username, password)
	rec := send(router, "POST", "/login", "", body)
	if rec.Code != http.StatusOK {
		t.Fatalf("login for %s failed with status %d", username, rec.Code)
	}

	var result struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatalf("login response was not valid JSON: %v", err)
	}
	return result.Token
}

const validEmployee = `{"name":"Anita","email":"anita@example.com","age":28,"salary":45000}`

func TestLoginSuccess(t *testing.T) {
	router := newTestRouter(t)
	token := loginAs(t, router, "alice", "adminpass")
	if token == "" {
		t.Error("expected a token, got an empty string")
	}
}

func TestLoginWrongPassword(t *testing.T) {
	router := newTestRouter(t)
	rec := send(router, "POST", "/login", "", `{"username":"alice","password":"wrong"}`)
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}

func TestProtectedRouteWithoutToken(t *testing.T) {
	router := newTestRouter(t)
	rec := send(router, "GET", "/api/profile", "", "")
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}

func TestProtectedRouteWithBadToken(t *testing.T) {
	router := newTestRouter(t)
	rec := send(router, "GET", "/api/profile", "not-a-real-token", "")
	if rec.Code != http.StatusUnauthorized {
		t.Errorf("expected 401, got %d", rec.Code)
	}
}

func TestProfileWithToken(t *testing.T) {
	router := newTestRouter(t)
	token := loginAs(t, router, "bob", "userpass")

	rec := send(router, "GET", "/api/profile", token, "")
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "bob") {
		t.Errorf("expected the profile to mention bob, got %s", rec.Body.String())
	}
}

func TestUserCannotCreateEmployee(t *testing.T) {
	router := newTestRouter(t)
	token := loginAs(t, router, "bob", "userpass")

	rec := send(router, "POST", "/api/employees", token, validEmployee)
	if rec.Code != http.StatusForbidden {
		t.Errorf("expected 403, got %d", rec.Code)
	}
}

func TestAdminCanCreateEmployee(t *testing.T) {
	router := newTestRouter(t)
	token := loginAs(t, router, "alice", "adminpass")

	rec := send(router, "POST", "/api/employees", token, validEmployee)
	if rec.Code != http.StatusCreated {
		t.Errorf("expected 201, got %d", rec.Code)
	}
}

func TestCreateEmployeeValidation(t *testing.T) {
	router := newTestRouter(t)
	token := loginAs(t, router, "alice", "adminpass")

	rec := send(router, "POST", "/api/employees", token, `{"name":"A","email":"not-an-email","age":10,"salary":-1}`)
	if rec.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rec.Code)
	}
}

func TestPagination(t *testing.T) {
	router := newTestRouter(t)
	token := loginAs(t, router, "alice", "adminpass")

	for i := 0; i < 5; i++ {
		send(router, "POST", "/api/employees", token, validEmployee)
	}

	rec := send(router, "GET", "/api/employees?page=2&limit=2", token, "")
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var result struct {
		Total      int        `json:"total"`
		TotalPages int        `json:"total_pages"`
		Data       []Employee `json:"data"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatalf("response was not valid JSON: %v", err)
	}

	if result.Total != 5 {
		t.Errorf("expected total 5, got %d", result.Total)
	}
	if result.TotalPages != 3 {
		t.Errorf("expected 3 pages, got %d", result.TotalPages)
	}
	if len(result.Data) != 2 {
		t.Errorf("expected 2 employees on page 2, got %d", len(result.Data))
	}
}

func TestDeleteEmployee(t *testing.T) {
	router := newTestRouter(t)
	token := loginAs(t, router, "alice", "adminpass")
	send(router, "POST", "/api/employees", token, validEmployee)

	rec := send(router, "DELETE", "/api/employees/1", token, "")
	if rec.Code != http.StatusOK {
		t.Errorf("expected 200 on first delete, got %d", rec.Code)
	}

	rec = send(router, "DELETE", "/api/employees/1", token, "")
	if rec.Code != http.StatusNotFound {
		t.Errorf("expected 404 on second delete, got %d", rec.Code)
	}
}

func TestRequestIDHeader(t *testing.T) {
	router := newTestRouter(t)
	rec := send(router, "POST", "/login", "", `{"username":"alice","password":"wrong"}`)
	if rec.Header().Get("X-Request-ID") == "" {
		t.Error("expected an X-Request-ID header on the response")
	}
}

func TestRecoveryMiddleware(t *testing.T) {
	router := newTestRouter(t)
	router.GET("/boom", func(c *gin.Context) {
		panic("something broke")
	})

	rec := send(router, "GET", "/boom", "", "")
	if rec.Code != http.StatusInternalServerError {
		t.Errorf("expected 500, got %d", rec.Code)
	}
}
