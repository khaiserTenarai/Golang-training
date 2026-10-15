package controller

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"secure-banking-api/model"

	"github.com/gin-gonic/gin"
)

type mockAuthService struct{}

func (mockAuthService) Register(ctx context.Context, req model.RegisterRequest) (*model.User, error) {
	return &model.User{ID: 1, Name: req.Name, Email: req.Email, Role: model.RoleUser}, nil
}

func (mockAuthService) Login(ctx context.Context, req model.LoginRequest) (string, error) {
	if req.Password != "correct-password" {
		return "", model.ErrInvalidCredentials
	}
	return "token", nil
}

func newAuthRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	ctl := NewAuthController(mockAuthService{})
	r := gin.New()
	r.POST("/register", ctl.Register)
	r.POST("/login", ctl.Login)
	return r
}

func post(r http.Handler, path, body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestRegisterValidation(t *testing.T) {
	r := newAuthRouter()
	tests := []struct {
		name string
		body string
		want int
	}{
		{"valid", `{"name":"Asha","email":"asha@example.com","password":"password123"}`, http.StatusCreated},
		{"bad email", `{"name":"Asha","email":"nope","password":"password123"}`, http.StatusBadRequest},
		{"short password", `{"name":"Asha","email":"asha@example.com","password":"123"}`, http.StatusBadRequest},
		{"missing name", `{"email":"asha@example.com","password":"password123"}`, http.StatusBadRequest},
		{"malformed json", `{`, http.StatusBadRequest},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := post(r, "/register", tc.body).Code; got != tc.want {
				t.Fatalf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestLogin(t *testing.T) {
	r := newAuthRouter()
	if got := post(r, "/login", `{"email":"a@b.com","password":"correct-password"}`).Code; got != http.StatusOK {
		t.Fatalf("want 200, got %d", got)
	}
	if got := post(r, "/login", `{"email":"a@b.com","password":"wrong"}`).Code; got != http.StatusUnauthorized {
		t.Fatalf("want 401, got %d", got)
	}
}
