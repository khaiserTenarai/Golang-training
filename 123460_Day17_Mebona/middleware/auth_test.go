package middleware

import (
	"employee-management/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func newRouter(ts *auth.TokenService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	ok := func(c *gin.Context) { c.String(http.StatusOK, "ok") }
	r.GET("/private", Auth(ts), ok)
	r.POST("/admin", Auth(ts), RequireRole(auth.RoleAdmin), ok)
	return r
}

// call sends one fake request and returns the HTTP status code.
func call(r *gin.Engine, method, path, token string) int {
	req := httptest.NewRequest(method, path, nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w.Code
}

func TestAuthMiddleware(t *testing.T) {
	ts := auth.NewTokenService("test-secret", time.Minute)
	r := newRouter(ts)
	adminToken, _ := ts.Generate(1, "admin", auth.RoleAdmin)
	userToken, _ := ts.Generate(2, "user", auth.RoleUser)

	tests := []struct {
		name   string
		method string
		path   string
		token  string
		want   int
	}{
		{"no token is rejected", "GET", "/private", "", 401},
		{"garbage token is rejected", "GET", "/private", "abc.def.ghi", 401},
		{"user can read", "GET", "/private", userToken, 200},
		{"user cannot use admin route", "POST", "/admin", userToken, 403},
		{"admin can use admin route", "POST", "/admin", adminToken, 200},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := call(r, tc.method, tc.path, tc.token); got != tc.want {
				t.Fatalf("expected %d, got %d", tc.want, got)
			}
		})
	}
}
