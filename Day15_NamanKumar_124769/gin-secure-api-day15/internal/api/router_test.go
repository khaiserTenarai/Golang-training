package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"gin-secure-api-day15/internal/config"
	"gin-secure-api-day15/internal/store"
)

func testRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return NewRouter(config.Config{
		Address: ":8080", JWTSecret: "test-secret", TokenTTLHours: 2,
		AllowedOrigins: []string{"http://localhost:3000"}, Mode: "test",
	}, store.NewMemoryStore())
}

func tokenFor(t *testing.T, router http.Handler, role string) string {
	t.Helper()
	body := `{"user_id":7,"role":"` + role + `"}`
	req := httptest.NewRequest(http.MethodPost, "/api/auth/token", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("token endpoint status = %d, body = %s", rec.Code, rec.Body.String())
	}
	var response map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	value, ok := response["access_token"].(string)
	if !ok || value == "" {
		t.Fatal("access_token missing")
	}
	return value
}

func TestHealthAndRequestID(t *testing.T) {
	router := testRouter()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set("X-Request-ID", "test-request-123")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if rec.Header().Get("X-Request-ID") != "test-request-123" {
		t.Fatalf("request ID header not preserved: %q", rec.Header().Get("X-Request-ID"))
	}
	if rec.Header().Get("X-Response-Time") == "" {
		t.Fatal("response time header missing")
	}
}

func TestProtectedRouteRequiresToken(t *testing.T) {
	rec := httptest.NewRecorder()
	testRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/api/employees", nil))
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rec.Code)
	}
}

func TestRoleAuthorization(t *testing.T) {
	router := testRouter()
	token := tokenFor(t, router, "employee")
	req := httptest.NewRequest(http.MethodDelete, "/api/employees/1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestCreateEmployeeValidation(t *testing.T) {
	router := testRouter()
	token := tokenFor(t, router, "admin")
	req := httptest.NewRequest(http.MethodPost, "/api/employees", bytes.NewBufferString(`{"name":"x","email":"bad","department":"","role":"root"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestPagination(t *testing.T) {
	router := testRouter()
	token := tokenFor(t, router, "employee")
	req := httptest.NewRequest(http.MethodGet, "/api/employees?page=1&limit=1", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var response struct {
		Data []any `json:"data"`
		Pagination struct {
			Page  int `json:"page"`
			Limit int `json:"limit"`
			Total int `json:"total"`
		} `json:"pagination"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if len(response.Data) != 1 || response.Pagination.Limit != 1 || response.Pagination.Total != 2 {
		t.Fatalf("unexpected pagination response: %+v", response)
	}
}

func TestRecoveryMiddleware(t *testing.T) {
	rec := httptest.NewRecorder()
	testRouter().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/demo/panic", nil))
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("expected 500, got %d", rec.Code)
	}
}
