package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"employee-management/auth"
	"employee-management/middleware"

	"github.com/gin-gonic/gin"
)

func TestHealth(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}
}

func TestRequestID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(middleware.RequestID())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"requestID": c.GetString("requestID"),
		})
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	if rec.Header().Get("X-Request-ID") == "" {
		t.Fatal("expected X-Request-ID header")
	}
}

func TestAuthRejectsMissingToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authService := auth.New(
		"test-secret",
		60,
		"admin",
		auth.HashPassword("admin123"),
		"user",
		auth.HashPassword("user123"),
	)

	r := gin.New()
	r.GET("/protected",
		middleware.RequireAuth(authService),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "authorized"})
		},
	)

	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status 401, got %d", rec.Code)
	}
}

func TestAdminRoleAllowed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authService := auth.New(
		"test-secret",
		60,
		"admin",
		auth.HashPassword("admin123"),
		"user",
		auth.HashPassword("user123"),
	)

	token, err := authService.Login("admin", "admin123")
	if err != nil {
		t.Fatal(err)
	}

	r := gin.New()
	r.GET("/admin",
		middleware.RequireAuth(authService),
		middleware.RequireRole("admin"),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "admin allowed"})
		},
	)

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", rec.Code, rec.Body.String())
	}
}

func TestUserDeniedAdminRole(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authService := auth.New(
		"test-secret",
		60,
		"admin",
		auth.HashPassword("admin123"),
		"user",
		auth.HashPassword("user123"),
	)

	token, err := authService.Login("user", "user123")
	if err != nil {
		t.Fatal(err)
	}

	r := gin.New()
	r.GET("/admin",
		middleware.RequireAuth(authService),
		middleware.RequireRole("admin"),
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "admin allowed"})
		},
	)

	req := httptest.NewRequest(http.MethodGet, "/admin", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("expected status 403, got %d", rec.Code)
	}
}
