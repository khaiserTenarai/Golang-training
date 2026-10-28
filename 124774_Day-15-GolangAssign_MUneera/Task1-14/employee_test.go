package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"employee-management/controller"
	"employee-management/middleware"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	r.Use(middleware.RequestID())

	// The service is nil because these tests don't call database handlers.
	ctl := controller.NewEmployeeController(
		nil,
		"test-secret-for-tests",
		60,
	)

	ctl.RegisterRoutes(r)

	return r
}

func TestHealth(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/healthz",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"want status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}
}

func TestProtectedRouteRejectsMissingToken(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/employees",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf(
			"want status %d, got %d",
			http.StatusUnauthorized,
			rec.Code,
		)
	}
}

func TestRequestIDHeader(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/healthz",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	requestID := rec.Header().Get("X-Request-ID")

	if requestID == "" {
		t.Fatal("expected X-Request-ID header")
	}
}

func TestAbortMiddleware(t *testing.T) {
	r := setupTestRouter()

	req := httptest.NewRequest(
		http.MethodGet,
		"/demo/abort",
		nil,
	)

	rec := httptest.NewRecorder()

	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf(
			"want status %d, got %d",
			http.StatusForbidden,
			rec.Code,
		)
	}
}
