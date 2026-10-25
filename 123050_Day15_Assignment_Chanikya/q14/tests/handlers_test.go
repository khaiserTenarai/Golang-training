package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"ems/auth"
	"ems/controller"
	"ems/middleware"

	"github.com/gin-gonic/gin"
)

func setupRouter(t *testing.T) *gin.Engine {
	t.Helper()
	gin.SetMode(gin.TestMode)

	authService, err := auth.New(
		strings.Repeat("x", 40),
		60,
		"admin", "Admin@12345",
		"user", "User@12345",
	)
	if err != nil {
		t.Fatal(err)
	}

	authController := controller.NewAuthController(authService)

	router := gin.New()
	router.POST("/login", authController.Login)

	protected := router.Group("/")
	protected.Use(middleware.RequireAuth(authService))

	protected.GET("/employees", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "employee list"})
	})

	admin := protected.Group("/")
	admin.Use(middleware.RequireRole("admin"))

	admin.POST("/employees", func(c *gin.Context) {
		c.JSON(201, gin.H{"message": "employee created"})
	})

	return router
}

func login(t *testing.T, router *gin.Engine, username, password string) string {
	t.Helper()

	body, _ := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})

	req := httptest.NewRequest(
		http.MethodPost, "/login", bytes.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("login failed: status=%d body=%s",
			rec.Code, rec.Body.String())
	}

	var response struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.Token == "" {
		t.Fatal("expected JWT token")
	}

	return response.Token
}

func TestLoginSuccess(t *testing.T) {
	router := setupRouter(t)
	login(t, router, "admin", "Admin@12345")
}

func TestLoginRejectsWrongPassword(t *testing.T) {
	router := setupRouter(t)

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		strings.NewReader(
			`{"username":"admin","password":"wrong"}`,
		),
	)
	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rec.Code)
	}
}

func TestProtectedRouteRejectsMissingToken(t *testing.T) {
	router := setupRouter(t)

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rec.Code)
	}
}

func TestValidTokenCanReadEmployees(t *testing.T) {
	router := setupRouter(t)
	token := login(t, router, "user", "User@12345")

	req := httptest.NewRequest(http.MethodGet, "/employees", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
}

func TestUserCannotCreateEmployee(t *testing.T) {
	router := setupRouter(t)
	token := login(t, router, "user", "User@12345")

	req := httptest.NewRequest(
		http.MethodPost,
		"/employees",
		strings.NewReader(`{}`),
	)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected 403, got %d", rec.Code)
	}
}

func TestAdminCanCreateEmployeeRoute(t *testing.T) {
	router := setupRouter(t)
	token := login(t, router, "admin", "Admin@12345")

	req := httptest.NewRequest(
		http.MethodPost,
		"/employees",
		strings.NewReader(`{}`),
	)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", rec.Code)
	}
}
