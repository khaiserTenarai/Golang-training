package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	jwtSecret = []byte("test-secret-key-at-least-32-characters-long")
	os.Exit(m.Run())
}

func resetData() {
	mu.Lock()
	defer mu.Unlock()

	users = make(map[string]User)
	tasks = make(map[int]Task)
	nextTaskID = 1
}

func request(
	router http.Handler,
	method, path string,
	body []byte,
	token string,
) *httptest.ResponseRecorder {
	req := httptest.NewRequest(method, path, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)
	return rec
}

func registerTestUser(t *testing.T, router http.Handler) {
	t.Helper()

	body := []byte(`{
		"email":"test@example.com",
		"password":"Password123!"
	}`)

	rec := request(router, http.MethodPost, "/register", body, "")
	if rec.Code != http.StatusCreated {
		t.Fatalf("registration: expected 201, got %d: %s",
			rec.Code, rec.Body.String())
	}
}

func loginTestUser(t *testing.T, router http.Handler) string {
	t.Helper()

	body := []byte(`{
		"email":"test@example.com",
		"password":"Password123!"
	}`)

	rec := request(router, http.MethodPost, "/login", body, "")
	if rec.Code != http.StatusOK {
		t.Fatalf("login: expected 200, got %d: %s",
			rec.Code, rec.Body.String())
	}

	var response map[string]string
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	token := response["token"]
	if token == "" {
		t.Fatal("expected a JWT token")
	}

	return token
}

func TestRegisterAndLogin(t *testing.T) {
	resetData()
	router := SetupRouter()

	registerTestUser(t, router)
	token := loginTestUser(t, router)

	if token == "" {
		t.Fatal("expected a token")
	}
}

func TestProtectedRouteRequiresJWT(t *testing.T) {
	resetData()
	router := SetupRouter()

	rec := request(router, http.MethodGet, "/tasks", nil, "")
	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rec.Code)
	}
}

func TestInvalidRegistration(t *testing.T) {
	resetData()
	router := SetupRouter()

	body := []byte(`{
		"email":"not-an-email",
		"password":"123"
	}`)

	rec := request(router, http.MethodPost, "/register", body, "")
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rec.Code)
	}
}

func TestUserCannotAccessAdminRoute(t *testing.T) {
	resetData()
	router := SetupRouter()

	registerTestUser(t, router)
	token := loginTestUser(t, router)

	rec := request(router, http.MethodGet, "/admin/users", nil, token)
	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d: %s",
			rec.Code, rec.Body.String())
	}
}

func TestTaskPagination(t *testing.T) {
	resetData()
	router := SetupRouter()

	registerTestUser(t, router)
	token := loginTestUser(t, router)

	for i := 1; i <= 3; i++ {
		body := []byte(`{"title":"Task"}`)
		rec := request(router, http.MethodPost, "/tasks", body, token)

		if rec.Code != http.StatusCreated {
			t.Fatalf("create task: expected 201, got %d", rec.Code)
		}
	}

	rec := request(router, http.MethodGet,
		"/tasks?page=1&limit=2", nil, token)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}

	var response struct {
		Data  []Task `json:"data"`
		Page  int    `json:"page"`
		Limit int    `json:"limit"`
		Total int    `json:"total"`
	}

	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}

	if len(response.Data) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(response.Data))
	}
	if response.Total != 3 {
		t.Fatalf("expected total 3, got %d", response.Total)
	}
	if response.Page != 1 || response.Limit != 2 {
		t.Fatalf("unexpected pagination metadata: %+v", response)
	}
}
